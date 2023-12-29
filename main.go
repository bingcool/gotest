package main

import (
	"fmt"
	"goTest/domain/System"
	"runtime"
	"time"
)

//var threadProfile = pprof.Lookup("threadcreate")

func main() {
	test()

	//System.GetRootDir()
	//err := Cmd.Execute()
	//if err != nil {
	//	log.Fatal("启动错误")
	//}
}

func init() {
	fmt.Println("Main package another init")
}

func test() {
	numCPU := runtime.NumCPU()
	fmt.Println("Number of CPUs:", numCPU)

	// 设置线程数量为 4
	runtime.GOMAXPROCS(1)

	// 信号监听
	System.EventLoopSigtermSignal()

	fmt.Println("start start ")
	//debug.SetMaxThreads(10)

	isErrDoneChan := make(chan int, 1)
	go func() {
		time.Sleep(10 * time.Hour)
		isErrDoneChan <- 1
	}()

	select {
	case <-isErrDoneChan:
		//发送消息给管理员
		fmt.Println("hhhhh")
	}
	time.Sleep(1 * time.Second)
}
