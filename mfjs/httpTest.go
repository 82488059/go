package main

import (
	"fmt"
	"net/url"
)

func main1() {
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
