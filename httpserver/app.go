package main

import (
	"fmt"
	"net/http"

	"httpserver/idv"
	"httpserver/logs"
	"httpserver/metadb"
	"httpserver/metaredis"
	"httpserver/src"

	"httpserver/sio"

	"github.com/gin-gonic/gin"

	//"github.com/go-redis/redis"

	socketio "github.com/googollee/go-socket.io"
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {
	logs.GetIns()
	logs.Trace("app start!")
	// 初始化
	src.Test()
	// 初始化redis连接
	metaredis.GetIns()
	// 初始化数据库连接
	metadb.GetIns()

	gp, er := metaredis.GetIns().Group()
	if er == 0 {
		fmt.Println("gp ", gp)
	}
	router := gin.Default()
	sioServer := socketio.NewServer(nil)

	sioServer.OnConnect("/", sio.OnConnect)
	sioServer.OnDisconnect("/", sio.OnDisconnect)

	sioServer.OnEvent("/host", "authorize", sio.AuthorizeEvent)
	sioServer.OnEvent("/host", "commit", sio.CommitEvent)

	sioServer.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	go sioServer.Serve()
	defer sioServer.Close()

	router.Use(GinMiddleware("http://localhost:8000"))
	// socket io
	router.GET("/socket.io/*any", gin.WrapH(sioServer))
	router.POST("/socket.io/*any", gin.WrapH(sioServer))
	// http
	router.StaticFS("/public", http.Dir("../asset"))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	router.GET("/pub/hi", idv.GetPubHi)
	router.GET("/pub/prtcfg", idv.GetPubPrtcfg)
	router.GET("/pub/lh/disklist", idv.GetPubLhDisklist)

	router.GET("/host/pub/hi", idv.GetPubHi)
	router.GET("/host/pub/prtcfg", idv.GetPubPrtcfg)
	router.GET("/host/pub/lh/disklist", idv.GetPubLhDisklist)

	router.POST("/host/pub/updimg", idv.PostPubUpdimg)
	router.POST("/host/pub/trial", idv.PostPubTrial)
	router.POST("/host/pub/join", idv.PostPubJoin)
	router.POST("/host/pub/hi", idv.PostPubHi)
	router.POST("/host/pub/openrooms", idv.PostPubOpenrooms)

	router.POST("/pub/updimg", idv.PostPubUpdimg)
	router.POST("/pub/trial", idv.PostPubTrial)
	router.POST("/pub/join", idv.PostPubJoin)
	router.POST("/pub/hi", idv.PostPubHi)
	router.POST("/pub/openrooms", idv.PostPubOpenrooms)

	router.Run("0.0.0.0:8000")
}
