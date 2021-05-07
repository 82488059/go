package param

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mfjs/random"
	"net/url"
	"strconv"
)

// DoZlibCompress 进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w, _ := zlib.NewWriterLevel(&in, zlib.BestCompression)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

// DoZlibUnCompress 进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

// Encode make post body
// {"mzsg":"do=GetContributePoints&v=8376&phpp=Android&phpl=ZH_CN&pvc=3.1.2.10356&pvb=2018-05-15+17%3a14%3a37","friendUid":"49424","pvpNewVersion":"1","OpenCardChip":"1"}
func Encode(mzsg string, kv map[string]interface{}) string {
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["mzsg"] = mzsg
	for k, v := range kv {
		data[k] = v
	}
	data["pvpNewVersion"] = "1"
	data["OpenCardChip"] = "1"
	text, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	//fmt.Println(text)
	bz := DoZlibCompress(text)
	bs := base64.StdEncoding.EncodeToString(bz)
	// fmt.Println(b4)
	zb := MakeZB(bs)
	fmt.Println(zb)
	// url body
	body := zb.Encode()
	return body
}

// Decode s='eNodzbEKwjAYBOB3CdhFW/I3qa1CBsmgU3XRwUVaE21Ak580RlB8d6PLwX0c3JvcX+OVLIlyYq2DdDZ40z+C3jljw5hF0bB6nuGAKFZWeWfUr9zEcXOSbYbxLFgBRVkAZVXaxV6UFJqcVjlUU6gnrAOegtVkRi7eaKv2RqU/vuAlT4YRW/08aD8aZ5NDsi1qKzuv5GDwT58vRKAyew=='
func Decode(s string) string {
	// s = urllib.unquote(s)
	data, _ := base64.StdEncoding.DecodeString(s)
	ss := DoZlibUnCompress(data)
	fmt.Println(string(ss))
	return string(ss)
}

// DecodeToMap ...
func DecodeToMap(s string) map[string]interface{} {
	data, _ := base64.StdEncoding.DecodeString(s)
	ss := DoZlibUnCompress(data)
	fmt.Println(string(ss))
	str := string(ss)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err != nil {
		fmt.Println(err)
	}
	return dat
}

// MakeZB make post body param z, b
func MakeZB(text string) url.Values {
	rs5 := random.GetRandomStringaz09(5)
	n := random.Int(1, 9)
	z := rs5 + strconv.Itoa(n) + random.GetRandomStringaz09(n) + text
	salt := "AAAAGQAAABhAAAAawAAAHMAAkAHMAYQAmACEAKABfACkASAAZgAAAGMAAABuAAAAYQAAA$10|654|18|180|137|823|77|791|77|252|"
	b := makeMd5([]byte(z + salt))
	zb := url.Values{}
	zb.Add("z", z)
	zb.Add("b", b)
	return zb
}

func makeMd5(buf []byte) string {
	var md5Ctx = md5.New()
	md5Ctx.Write(buf)
	b := hex.EncodeToString(md5Ctx.Sum(nil))
	return b
}
