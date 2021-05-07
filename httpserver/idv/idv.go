package idv

import (
	"fmt"
	"httpserver/metadb"
	"httpserver/metaredis"

	"github.com/gin-gonic/gin"
)

func PostPubUpdimg(c *gin.Context) {
	fmt.Println("post/pub/updimg")

	c.JSON(200, gin.H{
		"message": "pong",
	})

}
func PostPubTrial(c *gin.Context) {
	metaredis.Test()
	// 初始化redis连接
	metaredis.GetIns()
	// 初始化数据库连接
	metadb.GetIns()

	group, errno := metaredis.GetIns().Group()
	if 0 == errno {
		fmt.Println("group ", group)
	}

	fmt.Println("post/pub/trial")
	//var r = io.ReadCloser
	mac := c.Query("mac")
	fmt.Println("mac:", mac)

	//resp := map[string]interface{}{}
	// type H map[string]interface{}
	resp := gin.H{}
	resp["svrid"] = "D050992EE3D8"
	resp["svrname"] = "idvsvr"
	// 测试组
	rooms := []map[string]interface{}{}
	// add room1
	rooms = append(rooms, map[string]interface{}{})
	rooms[0]["name"] = "testroom1"
	rooms[0]["id"] = 1
	rooms[0]["section"] = "room1"
	rooms[0]["joinid"] = 10
	// 机器1
	slots1 := []map[string]interface{}{}
	slots1 = append(slots1, map[string]interface{}{})
	slots1[0]["id"] = 10
	slots1[0]["name"] = "r1i0"
	slots1[0]["iplan"] = "192.168.1.165"
	rooms[0]["slots"] = slots1
	// 默认选中
	// room2
	rooms = append(rooms, map[string]interface{}{})
	rooms[1]["name"] = "testroom2"
	rooms[1]["id"] = 2
	rooms[1]["section"] = "room2"
	rooms[1]["joinid"] = 20
	// 机器1
	slots2 := []map[string]interface{}{}
	slots2 = append(slots2, map[string]interface{}{})
	slots2[0]["id"] = 20
	slots2[0]["name"] = "r2i0"
	slots2[0]["iplan"] = "192.168.1.166"
	// 机器2
	slots2 = append(slots2, map[string]interface{}{})
	slots2[1]["id"] = 21
	slots2[1]["name"] = "r2i1"
	slots2[1]["iplan"] = "192.168.1.167"
	// 机器3
	slots2 = append(slots2, map[string]interface{}{})
	slots2[2]["id"] = 22
	slots2[2]["name"] = "r2i2"
	slots2[2]["iplan"] = "192.168.1.168"
	// 机器4
	slots2 = append(slots2, map[string]interface{}{})
	slots2[3]["id"] = 23
	slots2[3]["name"] = "r2i3"
	slots2[3]["iplan"] = "192.168.1.169"
	// 机器5
	slots2 = append(slots2, map[string]interface{}{})
	slots2[4]["id"] = 24
	slots2[4]["name"] = "r2i4"
	slots2[4]["iplan"] = "192.168.1.170"
	// 机器6
	slots2 = append(slots2, map[string]interface{}{})
	slots2[5]["id"] = 25
	slots2[5]["name"] = "r2i5"
	slots2[5]["iplan"] = "192.168.1.171"

	rooms[1]["slots"] = slots2

	resp["rooms"] = rooms
	resp["acctype"] = "ipaddr"
	resp["acchost"] = "192.168.1.4"
	resp["accport"] = "8000"
	resp["accproto"] = "http"

	c.JSON(200, resp)

}
func PostPubJoin(c *gin.Context) {
	fmt.Println("post/pub/join")

	c.JSON(200, gin.H{
		"message": "pong",
	})

}
func PostPubHi(c *gin.Context) {
	fmt.Println("post/pub/hi")

	c.JSON(200, gin.H{
		"message": "pong",
	})

}
func PostPubOpenrooms(c *gin.Context) {
	fmt.Println("post/pub/openrooms")

	c.JSON(200, gin.H{
		"message": "pong",
	})

}

func GetPubHi(c *gin.Context) {
	fmt.Println("get/pub/hi")

	c.JSON(200, gin.H{
		"message": "pong",
	})

}
func GetPubPrtcfg(c *gin.Context) {
	fmt.Println("get/pub/prtcfg")

	c.JSON(200, gin.H{
		"message": "pong",
	})

}
func GetPubLhDisklist(c *gin.Context) {
	fmt.Println("get/pub/lh/disklist")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
