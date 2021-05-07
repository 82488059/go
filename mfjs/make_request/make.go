package main

import (
	"fmt"
)

type make struct {
	num int
}

// makeMfjsURL ...
func (mk *make) makeMfjsURL(uri string, do string) string {
	mk.num++
	url := ""
	if do == "" {
		url = fmt.Sprintf("%s?&v=%d&phpp=Android&phpl=ZH_CN&pvc=3.1.1&pvb=2018-02-13+16%%3a02%%3a48", uri, mk.num)
	} else {
		url = fmt.Sprintf("%s?do=%s&v=%d&phpp=Android&phpl=ZH_CN&pvc=3.1.1&pvb=2018-02-13+16%%3a02%%3a48", uri, do, mk.num)
	}
	return url
}
