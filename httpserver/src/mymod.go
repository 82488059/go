package src

import (
	"encoding/json"
	"fmt"
	"log"

	goredis "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson"
)

func Test() {
	// Redigo Client  "github.com/gomodule/redigo/redis"
	rh := rejson.NewReJSONHandler()
	// connect redis
	conn, err := redis.Dial("tcp", "192.168.1.131:6379")
	if err != nil {
		return
	}
	// 绑定
	rh.SetRedigoClient(conn)
	js1, err := redis.Bytes(rh.JSONGet("group", "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return
	}
	// group
	var group interface{}
	err = json.Unmarshal(js1, &group)
	if err != nil {
		return
	}
	fmt.Println("group1")
	fmt.Println(group)
	// _1
	js1, err = redis.Bytes(rh.JSONGet("group", "_1"))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return
	}
	var _1 interface{}
	err = json.Unmarshal(js1, &_1)
	if err != nil {
		return
	}
	fmt.Println("_1")
	fmt.Println(_1)

	err = conn.Close()
	if err != nil {
		return
	}

	// GoRedis Client "github.com/go-redis/redis"
	// connect redis
	cli := goredis.NewClient(&goredis.Options{Addr: "192.168.1.131:6379", // use default Addr
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 绑定
	rh.SetGoRedisClient(cli)
	js2, err := redis.Bytes(rh.JSONGet("group", "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return
	}
	//
	var group2 interface{}
	err = json.Unmarshal(js2, &group2)
	if err != nil {
		return
	}
	fmt.Println("group2")
	fmt.Println(group2)
	err = cli.Close()
	if err != nil {
		return
	}
	return
}
