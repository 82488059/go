package metadb

import (
	"fmt"
	"sync"
)

type singletonDb struct {
}

var ins *singletonDb
var once sync.Once

func GetIns() *singletonDb {
	once.Do(func() {
		ins = &singletonDb{}
	})
	return ins
}

func Test() {
	fmt.Println("redis.go")
}
