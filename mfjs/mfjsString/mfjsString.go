package mfjsString
import (
	"fmt"
	"mfjs/random"
)
//
var Udid = "08:00:27:24:CB:9B"

//
var MfjsCookie = "_sid=5iubge64hkkohr47v6hi48hjb6; expires=Sat, 28-Apr-2018 14:28:01 GMT; path=/"

//
var GameName = "ANQUCARD-DROID-CHS-YUZHUANG02"

//
var MainUrl = "/pp/httpService.do"

// MainHost
var MainHost = "bj.muhepp.com"

// MainHeader ...
var MainHeaders map[string]string = map[string]string{"Host": "bj.muhepp.com",
	"User-Agent":      "UnityPlayer/5.6.3f1 (http://unity3d.com)",
	"Accept":          "*/*",
	"Accept-Encoding": "identity",
	"Content-Type":    "Application/json",
	"X-Unity-Version": "5.6.3f1"}

// MainHeader
func MainHeader() map[string]string {

	return MainHeaders
}

// MakeMainHeader ...
func MakeMainHeader() map[string]string {
	MainHeader := map[string]string{"Host": "bj.muhepp.com",
		"User-Agent":      "UnityPlayer/5.6.3f1 (http://unity3d.com)",
		"Accept":          "*/*",
		"Accept-Encoding": "identity",
		"Content-Type":    "Application/json",
		"X-Unity-Version": "5.6.3f1"}

	return MainHeader
}

func mfjs_user_cookie(){
    cookie = "_sid="
    for x =0; x < 13; x++ {
		cookie += fmt.Sprintf(":%x", random.Int(0, 255))
	}
    return cookie + "; expires=Sat, 28-Apr-2018 14:28:01 GMT; path=/"
}

func Make_mfjs_header(host) map[string]string {
	cookie := mfjs_user_cookie()
	header = map[string]string

    header = {    "Host": host,
                  "Cookie": cookie,
                  "Accept": "text/xml, application/xml, application/xhtml+xml, text/html;q=0.9, text/plain;q=0.8, text/css, "
                            "image/png, image/jpeg, image/gif;q=0.8, application/x-shockwave-flash, video/mp4;q=0.9, "
                            "flv-application/octet-stream;q=0.8, video/x-flv;q=0.7, audio/mp4, application/futuresplash, "
                            "*/*;q=0.5",
                  "User-Agent": "Mozilla/5.0 (Android; U; zh-CN) AppleWebKit/533.19.4 (KHTML, like Gecko) AdobeAIR/23.0",
                  "x-flash-version": "23,0,0,162",
				  "Connection": "Keep-Alive",
				  "Cache-Control": "no-cache",
                  "Referer": "app:/assets/CardMain.swf",
                  "Content-Type": "application/x-www-form-urlencoded"}
    return header
}