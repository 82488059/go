package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Connect struct {
	proxy string
}

// Post ...
func Post(myurl string, data url.Values) int {
	resp, err := http.PostForm(myurl, data)
	if err != nil {
		// handle error
		return -1
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		return -1
	}

	fmt.Println(string(body))

	return 0
}

// ProxyPost  post ...
func ProxyPost(myurl string, data url.Values, proxyAddress string) int {
	urlParse, _ := url.Parse(proxyAddress)
	proxy := http.ProxyURL(urlParse)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxy,
		},
	}
	reqest, err := client.Head(myurl)
	reqest.Header.Add("Cookie", "xxxxxx")
	reqest.Header.Add("User-Agent", "xxx")
	reqest.Header.Add("X-Requested-With", "xxxx")
	if err != nil {
		// handle error
		return -1
	}
	resp, err := client.PostForm(myurl, data)
	// resp, err := http.PostForm(myurl, data)
	if err != nil {
		// handle error
		return -1
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		return -1
	}

	fmt.Println(string(body))

	return 0
}

/*
func main() {
	data := url.Values{}
	proxy := "159.65.47.178:8888"
	n := ProxyPost("http://www.baidu.com", data, proxy)
	fmt.Println(n)
}
*/
