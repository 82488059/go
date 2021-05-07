package main

import (
	"flag"
	"fmt"
	"goDemo/filecontrol"
	"goDemo/logger"
	"os"
	"os/exec"
)

var godaemon = flag.Bool("d", false, "run app as a daemon with -d=true or -d true.")

func init1() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *godaemon {
		fmt.Println(flag.Args())
		cmd := exec.Command(os.Args[0], flag.Args()[1:]...)
		cmd.Start()
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		*godaemon = false
		fmt.Println(flag.Args())
		os.Exit(0)
	}
}

func main() {
	err := logger.LogInit()
	if err != nil {
		fmt.Printf("parse log config err, please check log.config!")
		return
	}
	filecontrol.StartFileControl()
	//阻塞等待退出
	wait := make(chan int)
	<-wait
	fmt.Println("goagent exit")
}
