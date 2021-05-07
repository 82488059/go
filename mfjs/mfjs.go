package main

//go mod init test3

import (
	"fmt"
	"mfjs/param"
	"mfjs/user"
	"net/url"
)

func testEncode() {
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["friendUid"] = "49424"
	//uri := "friend.php/"
	do := "do=GetContributePoints&v=8376&phpp=Android&phpl=ZH_CN&pvc=3.1.2.10356&pvb=2018-05-15+17%3a14%3a37"
	// make
	b := param.Encode(do, data)
	fmt.Println(b)
	// decode
	a, _ := url.ParseQuery(b)
	fmt.Println(a)
	js := param.DecodeToMap("eNodzbEKwjAYBOB3CdhFW/I3qa1CBsmgU3XRwUVaE21Ak580RlB8d6PLwX0c3JvcX+OVLIlyYq2DdDZ40z+C3jljw5hF0bB6nuGAKFZWeWfUr9zEcXOSbYbxLFgBRVkAZVXaxV6UFJqcVjlUU6gnrAOegtVkRi7eaKv2RqU/vuAlT4YRW/08aD8aZ5NDsi1qKzuv5GDwT58vRKAyew==")
	fmt.Println(js)

	// url encode
	v := url.Values{}
	v.Add("a", "aa")
	v.Add("b", "bb")
	v.Add("c", "有没有人")
	body := v.Encode()
	fmt.Println(v)
	fmt.Println(body)
	// url decode
	m, _ := url.ParseQuery(body)
	fmt.Println(m)
}

func main() {
	var u user.User
	u.Make("", "")
}
