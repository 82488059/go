package logs

import (
	"fmt"
	"log"
	"os"
	"sync"
)

const (
	LOG_ERROR = 1
	LOG_WRONG = 2
	LOG_INFO  = 3
	LOG_DEBUG = 5
	LOG_TRACE = 6
)

type singletonRedis struct {
	fileName string
	logger   *log.Logger
	logLevel int
}

type sh interface {
	Group() (map[string]interface{}, int)
}

var ins *singletonRedis
var once sync.Once

func GetIns() *singletonRedis {
	once.Do(func() {
		ins = &singletonRedis{}
		ins.fileName = "debug.log"
		logFile, err := os.OpenFile(ins.fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
		if err != nil {
			log.Fatalln("open file error !")
		} else {
			ins.logger = log.New(logFile, "logger: ", log.Llongfile)
		}
		ins.logLevel = LOG_TRACE
	})
	return ins
}

func Trace(s string, params ...interface{}) {
	ins = GetIns()
	if nil == ins {
		return
	}
	if ins.logLevel < LOG_TRACE {
		return
	}
	ins.logger.Printf(s, params...)
	fmt.Printf(s, params...)
	return
}

func Error(s string, params ...interface{}) {
	ins = GetIns()
	if nil == ins {
		return
	}
	if ins.logLevel < LOG_ERROR {
		return
	}
	ins.logger.Printf(s, params...)
	return
}

func Wrong(s string, params ...interface{}) {
	ins = GetIns()
	if nil == ins {
		return
	}
	if ins.logLevel < LOG_WRONG {
		return
	}
	ins.logger.Printf(s, params...)
	return
}

func Info(s string, params ...interface{}) {
	ins = GetIns()
	if nil == ins {
		return
	}
	if ins.logLevel < LOG_INFO {
		return
	}
	ins.logger.Printf(s, params...)
	return
}

func Debug(s string, params ...interface{}) {
	ins = GetIns()
	if nil == ins {
		return
	}
	if ins.logLevel < LOG_DEBUG {
		return
	}
	ins.logger.Printf(s, params...)
	return
}
