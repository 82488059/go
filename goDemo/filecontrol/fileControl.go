package filecontrol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"goDemo/logger"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"

	mmap "github.com/edsrzf/mmap-go"
	jsoniter "github.com/json-iterator/go"
)

var loggerFile = logger.NewPrefixLogger("fileControl")
var nKerMsgLen = 0

//StartUser 启动用户查询协程
func StartFileControl() {
	msgHandle()
	//startGetMsg()
	//startKernelUDP()
	//startGetKernelMsg()
}

func startGetMsg() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				loggerFile.Error("get策略信息失败", r)
				return
			}
		}()

		msgHandle()
		fileTimer := time.NewTimer(10 * time.Second)
		for {
			select {
			case <-fileTimer.C:
				msgHandle()
				fileTimer.Reset(10 * time.Second)
			}
		}
	}()
}

//请求网络策略消息，请求需要分发的时候直接写文件，然后通知内核读取消息
func msgHandle() {
	defer func() {
		if r := recover(); r != nil {
			loggerFile.Error("get策略信息失败", r)
			return
		}
	}()
	resp, err := http.Get("http://192.168.103.212:8003/api/bbdpoc")
	if err != nil || resp.StatusCode != http.StatusOK {
		loggerFile.Debug("Get请求失败", err)
		return
	}

	//接收到返回数据
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		loggerFile.Debug("读取Get请求Body错误", err)
		return
	}
	//loggerFile.Debug("请求到数据：", string(body[:]))

	var retMap []interface{}
	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsonIterator.Unmarshal([]byte(body), &retMap)
	if err != nil {
		loggerFile.Debug("json转换Map失败！", err)
		return
	}
	writeToFile(retMap)
}

func writeToFile(msg []interface{}) {
	defer func() {
		if r := recover(); r != nil {
			loggerFile.Error("get策略信息失败", r)
			return
		}
	}()
	//f, err := os.OpenFile(BfxPathFileControl, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	f, err := os.OpenFile("BfxPathFileControl.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}

	defer f.Close()
	for _, value := range msg {
		mapValue, ok := value.(map[string]interface{})
		if !ok {
			continue
		}
		var fileProtect FilePolicy
		uuid, err := getUUIDFromUserName(mapValue["username"].(string))
		if err != nil {
			loggerFile.Error("未查找到uuid", mapValue["username"].(string))
			continue
		}
		fileProtect.username = uuid
		copy(fileProtect.procpath[:], []byte(mapValue["procpath"].(string))[:])
		prochash := mapValue["prochash"].(string)
		prochashbyte, err := converHashToByte(prochash)
		if err != nil {
			loggerFile.Error("转换hash失败", prochash)
			continue
		}
		copy(fileProtect.prochash[:], prochashbyte[:])
		copy(fileProtect.filepath[:], []byte(mapValue["filepath"].(string))[:])
		fileProtect.mode = 0
		modeString := mapValue["mode"].(string)
		modeList := strings.Split(modeString, ",")
		for _, modeVaule := range modeList {
			mode, err := strconv.Atoi(modeVaule)
			if err != nil {
				continue
			}
			fileProtect.mode |= uint32(mode)
		}
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, fileProtect)
		f.Write(buf.Bytes())
	}
	fileInfo, err := f.Stat()
	if fileInfo.Size() > 0 {
		fc, err := os.Open("/dev/bfx")
		if err != nil {
			loggerFile.Debug("打开文件设备失败！")
			return
		}
		defer fc.Close()
		syscall.Syscall(syscall.SYS_IOCTL, fc.Fd(), NotifyPolicy, 0)
	}
}

func startKernelUDP() {
	var ker KerMsg
	nKerMsgLen = int(unsafe.Sizeof(ker))
	go func() {
		defer func() {
			if r := recover(); r != nil {
				loggerFile.Error("安全协程崩溃", r)
				return
			}
		}()
		udpAddr, err := net.ResolveUDPAddr("udp", ":10240")
		if err != nil {
			loggerFile.Error("绑定端口失败", err)
			return
		}

		conn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			loggerFile.Error("内核监听失败", err)
			return
		}
		defer conn.Close()

		for {
			recvUDPMsg(conn)
		}
	}()
}

func recvUDPMsg(conn *net.UDPConn) {
	var buf [1024]byte

	n, _, err := conn.ReadFromUDP(buf[0:])
	if err != nil || nKerMsgLen != n {
		loggerFile.Error("读取出错", err, n)
		return
	}
	var msgArr []interface{}
	handleKernelMsg(buf[0:n], &msgArr)
	sendKerMsg(&msgArr)
}

func startGetKernelMsg() {
	var ker KerMsg
	nKerMsgLen = int(unsafe.Sizeof(ker))
	go func() {
		defer func() {
			if r := recover(); r != nil {
				loggerFile.Error("处理内核消息崩溃", r)
				return
			}
		}()

		f, err := os.Open("/dev/bfx")
		if err != nil {
			loggerFile.Error("打开文件设备失败！")
			return
		}
		defer f.Close()
		f.SetReadDeadline(time.Time{})

		readByte := make([]byte, 1)
		for {
			nCount, err := f.Read(readByte)
			if nCount == 0 {
				continue
			}
			_ = err
			mem, err := mmap.MapRegion(f, nCount*1024*4, mmap.RDONLY, 0, 0)
			var msgArr []interface{}
			for i := 0; i < nCount; i++ {
				bytePage := mem[i*1024*4 : (i+1)*1024*4]
				handleKernelMsg(bytePage, &msgArr)
			}
			mem.Unmap()

			sendKerMsg(&msgArr)
		}
	}()
}

func handleKernelMsg(kerMsg []byte, msgArr *[]interface{}) {
	tempByte := make([]byte, nKerMsgLen, nKerMsgLen)

	copy(tempByte[:], kerMsg[0:nKerMsgLen])
	KerMsgStruct := (*KerMsg)(unsafe.Pointer(&(kerMsg)[0]))

	msgItem := make(map[string]interface{})
	msgItem["procpath"] = byteToStr(KerMsgStruct.procpath[:])
	msgItem["filepath"] = byteToStr(KerMsgStruct.filepath[:])
	var modeStringArr []string
	if KerMsgStruct.mode&Aram != 0 {
		modeStringArr = append(modeStringArr, "1")
	}

	if KerMsgStruct.mode&Read != 0 {
		modeStringArr = append(modeStringArr, "2")
	}

	if KerMsgStruct.mode&Write != 0 {
		modeStringArr = append(modeStringArr, "4")
	}

	if KerMsgStruct.mode&Delete != 0 {
		modeStringArr = append(modeStringArr, "8")
	}

	if KerMsgStruct.mode&Rename != 0 {
		modeStringArr = append(modeStringArr, "16")
	}

	if KerMsgStruct.mode&Access != 0 {
		modeStringArr = append(modeStringArr, "32")
	}

	modeString := strings.Join(modeStringArr, "、")

	msgItem["mode"] = modeString

	if KerMsgStruct.result == Allow {
		msgItem["result"] = "Allow"
	} else {
		msgItem["result"] = "Reject"
	}
	msgItem["createdTime"] = byteToStr(KerMsgStruct.createdTime[:])
	*msgArr = append(*msgArr, msgItem)
}

func sendKerMsg(msgArr *[]interface{}) {

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	retByte, err := jsonIterator.Marshal(msgArr)
	if err != nil {
		return
	}

	reqbody := byteToStr(retByte[:])
	//创建请求
	postReq, err := http.NewRequest("POST",
		"http://192.168.103.212:8003/api/bbdpoclog", //post链接
		strings.NewReader(reqbody))                  //post内容

	fmt.Println(reqbody)
	return

	if err != nil {
		fmt.Println("POST请求:创建请求失败", err)
		return
	}

	//增加header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	//执行请求
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		fmt.Println("POST请求:创建请求失败", err)
		return
	}
	defer resp.Body.Close()
	//读取响应
	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		fmt.Println("POST请求:读取body失败", err)
		return
	}
	fmt.Println("POST请求:创建成功", string(body))
}

func getUUIDFromUserName(userName string) (uid uint32, err error) {
	usernameTrime := strings.Trim(userName, " ")
	fp, err := os.Open("/etc/passwd")
	if err != nil {
		return 0, errors.New("打开passwd文件失败！")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		passwdInfo := scanner.Text()
		passwdInfoList := strings.Split(passwdInfo, ":")
		if len(passwdInfoList) != 7 {
			continue
		} else {
			if passwdInfoList[0] == usernameTrime {
				uuid, err := strconv.Atoi(passwdInfoList[2])
				if err != nil {
					return 0, nil
				}
				return uint32(uuid), nil
			}
		}
	}
	return 0, errors.New("未查找到UUID")
}

func converHashToByte(hashProc string) (ret []byte, err error) {
	if len(hashProc) != 32 {
		return nil, errors.New("hash字符串错误！")
	}

	var hashProcByte [16]byte
	for i := 0; i < 16; i++ {
		nRet, err := strconv.ParseInt(hashProc[2*i:(2*i+2)], 16, 32)
		if err != nil {
			return nil, errors.New("hash字符串错误！")
		}
		hashProcByte[i] = uint8(nRet)
	}
	return hashProcByte[:], nil
}

func byteToStr(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
