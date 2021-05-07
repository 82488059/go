package user

import (
	"mfjs/mfjsString"
)

// User ...
type User struct {
	name string
	pwd  string
}

// MakeUser ...
func MakeUser(name string, pwd string) User {
	var u User
	u.name = name
	u.pwd = pwd
	return u
}

// Make ...
func (u *User) Make(name string, pwd string) User {
	u.name = name
	u.pwd = pwd
	return *u
}

// Login ...
func (u *User) Login() int {
	// http.Post()
	return -1
}

func (u *User) mfjsMakeUdid() string {
	return ""
}

// loginOne ...
func loginOne(user_info *User) {
	body := make(map[string]interface{})
	// body["callPara"] = make(map[string]string)
	body["serviceName"] = "login"

	callPara := make(map[string]string)
	callPara["userPassword"] = ""
	callPara["userName"] = ""
	callPara["gameName"] = "CARDNEW-ANDROID-CHS"
	callPara["udid"] = user_info.mfjsMakeUdid()
	callPara["clientType"] = "android"
	callPara["releaseChannel"] = "android"
	callPara["locale"] = "ZH_CN"
	callPara["idfa"] = ""

	// udid := user_info.mfjsMakeUdid()
	//	body := "{\"callPara\":{\"userPassword\":\"" + user_info.pwd + "\"\",\"userName\":\"" + user_info.name
	//	+"\"\",\"gameName\":\"CARDNEW-ANDROID-CHS\",\"udid\":\"" + udid
	//	+"\",\"clientType\":\"android\",\"releaseChannel\":\"android\",\"locale\":\"ZH_CN\",\"idfa\":\"\"},\"serviceName\":\"login\"}"

	host := mfjsString.MainHost

	url := mfjsString.MainUrl
	header := mfjsString.MainHeaders

	user_info.MfjsHeader = mfjsString.Make_mfjs_header(user_info.Host)

	user_info.conn.mfjs_host = user_info.Host

	if conf.login_use_proxy {
		res_str = user_info.conn.post_service(host, url, body, header)
	} else {
		res_str = Connect.three_post_request(host, url, body, header)
	}
	return res_str
}
