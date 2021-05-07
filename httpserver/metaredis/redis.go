package metaredis

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson"
)

type singletonRedis struct {
	rh *rejson.Handler
}

type sh interface {
	Group() (map[string]interface{}, int)
}

var ins *singletonRedis
var once sync.Once

func GetIns() *singletonRedis {
	once.Do(func() {
		ins = &singletonRedis{}
		ins.rh = rejson.NewReJSONHandler()
		// connect redis
		conn, err := redis.Dial("tcp", "192.168.1.131:6379")
		if err != nil {
			ins.rh = nil
			ins = nil
			return
		}
		// 绑定
		ins.rh.SetRedigoClient(conn)

	})
	return ins
}

func Test() {
	fmt.Println("redis.go")
}

// sr *singletonRedis redis连接
func (sr *singletonRedis) Group() (map[string]interface{}, int) {
	groups := map[string]interface{}{}
	// group
	var group interface{}
	js1, err := redis.Bytes(sr.rh.JSONGet("group", "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return nil, -1
	}
	err = json.Unmarshal(js1, &group)
	if err != nil {
		return nil, -1
	}
	fmt.Println("group1")
	fmt.Println(group)
	// _1
	js1, err = redis.Bytes(sr.rh.JSONGet("group", "_1"))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return nil, -1
	}
	var _1 interface{}
	err = json.Unmarshal(js1, &_1)
	if err != nil {
		return nil, -1
	}
	fmt.Println("_1")
	fmt.Println(_1)

	return groups, 0
}

// sr *singletonRedis redis连接
func (sr *singletonRedis) Get(key string, key2 string) (interface{}, int) {
	// group
	var group interface{}
	js1, err := redis.Bytes(sr.rh.JSONGet(key, key2))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return nil, -1
	}
	err = json.Unmarshal(js1, &group)
	if err != nil {
		log.Fatalf("Failed to Unmarshal")
		return nil, -1
	}
	log.Fatalln("Failed to JSONGet")
	fmt.Println("group1")
	fmt.Println(group)
	return group, 0
}
