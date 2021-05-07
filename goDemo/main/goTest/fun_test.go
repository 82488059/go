package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"goagent/audit/common"
	"goagent/auxi"
	"goagent/common"
	"log"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"

	"database/sql"

	"github.com/edsrzf/mmap-go"
	_ "github.com/mattn/go-sqlite3"
)

func TestString(t *testing.T) {
	strtest := `CentOS release 6.9 (Final)
	Kernel \r on an \m
	
	`
	strbyte := []byte(strtest)
	for i := 0; i < len(strbyte); i++ {
		if strbyte[i] == ' ' && strbyte[i+1] != '\\' {
			strbyte[i] = '_'
		} else if strbyte[i] == '\n' || strbyte[i] == '\\' {
			strbyte[i] = 0
			strbyte = strbyte[:i]
			break
		}
	}
	<-time.NewTimer(1 * time.Second).C
}

func TestWorld(t *testing.T) {
	interarr, err := net.Interfaces()
	if err != nil {
		return
	}
	//mac地址
	for _, inter := range interarr {
		addrs, err := inter.Addrs()
		if err != nil {
			fmt.Println("获取IP地址失败！" + err.Error())
			return
		}
		//ip地址一个ip4一个ip6
		for _, addr := range addrs {
			fmt.Println(reflect.TypeOf(addr))
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				if "192.168.104.110" == ipnet.IP.String() {
					fmt.Println(inter.HardwareAddr.String())
					return
				}
			}
		}
	}
	return
}

func TestTranslate(t *testing.T) {
	type B struct {
		One [35]byte
		TWO int32
	}

	type A struct {
		// should be exported member when read back from buffer
		ss  B
		One [35]byte
		Two int32
	}

	var a A
	copy(a.ss.One[:], "xiong yan fei1")
	a.ss.TWO = 10
	copy(a.One[:], "xiong yan fei")
	a.Two = 5
	var aa A

	//fmt.Println(a)

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, a)
	if err != nil {
		fmt.Println(err)
	}

	testName := "123123"
	var testByte [16]byte
	copy(testByte[:], []byte(testName)[0:16])
	fmt.Println("%%%%%%%%%%%%%%%%%%%%", testByte)

	pPtr := (*A)(unsafe.Pointer(&(buf.Bytes())[0]))
	//fmt.Println(pPtr)
	fmt.Println("-------------", string(pPtr.ss.One[:]))
	// err = binary.Read(buf, binary.LittleEndian, &aa)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println(buf.Bytes())
	fmt.Println("after aa is ", aa)
}

func Test_Translate1(t *testing.T) {

}

func TestJson(t *testing.T) {
	var nTest uint32
	var nByte [4]byte

	nTest = 1

	binary.BigEndian.PutUint32(nByte[:], nTest)

	nTest2 := binary.LittleEndian.Uint32(nByte[:])
	fmt.Println(nTest2)
	fmt.Println(nByte)
	binary.LittleEndian.PutUint32(nByte[:], nTest)
	fmt.Println(nByte)

}

func TestChannel(t *testing.T) {
	msg := make(chan int, 10)

	work := func(name int) {
		for {
			select {
			case <-msg:
				fmt.Printf("%d", name)
			}
		}
	}

	go func() {
		i := 1000
		for {
			i--
			msg <- i
			if i <= 0 {
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		go work(i)
	}
	time.Sleep(50000 * time.Second)
}

func TestType(t *testing.T) {
	type test1 int
	type test2 int

	var s1 test1
	s1 = 10
	var s2 test2
	s2 = 10

	var obj interface{}
	obj = s1
	recongiseType := func() {
		switch obj.(type) {
		case test1:
			fmt.Printf("type1\n")
		case test2:
			fmt.Printf("type2\n")
		default:
			fmt.Println(1)
		}
	}
	recongiseType()
	obj = s2
	recongiseType()
}

func TestLen(t *testing.T) {
	var testArray [10]byte
	nIndex := strings.Index(string(testArray[:]), "0")
	fmt.Println(nIndex)

	index := 0
	for {
		testArray[index] = (byte)(index + 1)
		index++
		if index > 6 {
			break
		}
	}
	testArray[6] = 37
	byteString := func(p []byte) []byte {
		for i := 0; i < len(p); i++ {
			if p[i] == 0 {
				return p[0:i]
			}
		}
		return p
	}

	str := byteString(testArray[:])
	fmt.Println(str)
}

func TestCFunc(t *testing.T) {
	utmpArray, err := auxi.GetAllUtmp()
	if err != nil {
		return
	}
	for _, value := range utmpArray {
		t := time.Unix(int64(value.TVSec), 0)
		nt := t.Format("2006-01-02T15:04:05+08:00")
		fmt.Println(nt)
		fmt.Println(string(value.UTUser[:]))
		fmt.Println(string(value.UTHost[:]))
	}
}

func TestIOTA(t *testing.T) {
	fmt.Println(common.PolicyOffLine)
	fmt.Println(common.PolicyOnLine)
}

func TestChannelStruct(t *testing.T) {

	type A struct {
		name string
		age  uint32
	}

	msg := make(chan interface{}, 10)

	var testA A
	testA.name = "xiongyanfei"
	testA.age = 30
	msg <- testA

	var testB A
	testB.name = "xiongyanfei"
	testB.age = 30
	msg <- &testB

	work := func() {
		for {
			select {
			case value, _ := <-msg:
				if valueA, ok := value.(A); ok {
					fmt.Println(valueA)
					valueA.name = "xiong"
					valueA.age = 31
				} else if valueA, ok := value.(*A); ok {
					fmt.Println(valueA)
					valueA.name = "xiongyan"
					valueA.age = 32
				} else if valueA, ok := value.([]byte); ok {
					fmt.Println("go:", valueA)
					valueA[3] = 3
				}

			}
		}
	}

	go work()

	var testByte [256]byte
	testByte[0] = 1
	msg <- testByte[:]
	time.Sleep(2 * time.Second)
	fmt.Println("no go", testByte)

	time.Sleep(2 * time.Second)
	fmt.Println(testA)
	fmt.Println(testB)

	//结论channel 发送结构体的时候是数据的拷贝，发送结构体的指针的时候发送的是结构体的指针值,会修改掉原数据

	time.Sleep(50000 * time.Second)
}

func TestSwitchStruct(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()
	var nType byte
	nType = 3

	var msgStruct auditCommon.MsgCommon

	msgStruct.NType = auditCommon.MsgType(nType)
	fmt.Println(msgStruct)
}

func TestByte2String(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()
	var testByte [60]byte
	testByte[0] = 'a'
	testByte[1] = 'b'
	var stringType string
	stringType = string(testByte[0:3])
	fmt.Println(stringType)

	mapTest := make(map[string]interface{})

	mapTest["count"] = 1
	arrTest := make([]interface{}, 120)
	arrTest = arrTest[0:0]
	mapTest["array"] = arrTest
	item := make(map[string]interface{})
	item["item1"] = "test"
	item2 := make(map[string]interface{})
	item2["item2"] = "test2"
	arrTest = append(arrTest, item)
	arrTest = append(arrTest, item2)
	//arrTest = arrTest[0:0]
	mapTest["array"] = arrTest
	fmt.Println(mapTest)

}

func TestByteMask(t *testing.T) {
	mask := []byte{0xff, 0xff, 0xfc, 0x00}
	maskStr := ""
	maskStr += strconv.Itoa(int(mask[0])) + "."
	maskStr += strconv.Itoa(int(mask[1])) + "."
	maskStr += strconv.Itoa(int(mask[2])) + "."
	maskStr += strconv.Itoa(int(mask[3]))
	fmt.Println(maskStr)
}

func Test_mmap(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			auxi.MakePanicTrace(r)
			return
		}
	}()
	f, err := os.Open("/dev/via")
	_ = err
	nReadCount := make([]byte, 100)
	f.Read(nReadCount)
	mem, err := mmap.Map(f, mmap.COPY, 0)
	fmt.Println(mem)
	mem.Unmap()
}

func Test_Var(t *testing.T) {
	type user struct {
		lock sync.Mutex
		name string
		age  int
	}

	u := new(user)
	//var u user
	u.lock.Lock()
	u.name = "123"
	u.lock.Unlock()
	fmt.Println(u)
}

func Test_Sqlite(t *testing.T) {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func TestTime(t1 *testing.T) {
	t := time.Now()
	fmt.Println(t)

	fmt.Println(t.UTC())

	fmt.Println(t.UTC().Format(time.UnixDate))

	fmt.Println(t.Unix())

	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	fmt.Println(timestamp)
	timestamp = timestamp[:10]
	fmt.Println(timestamp)
}

func TestPrintln(t *testing.T) {
	strList := []string{"1", "2", "3"}
	for _, value := range strList {
		go func() {
			fmt.Println(value)
		}()
	}
}

func TestStruct(t *testing.T) {
	type Person struct {
	}

	type Techer struct {
		Person
	}

}
