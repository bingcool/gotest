package main

import (
	"fmt"
	"goTest/domain/System"
	"gopkg.in/yaml.v2"
	"runtime"
	"time"
)

//var threadProfile = pprof.Lookup("threadcreate")

func main() {
	//test1()

	//System.GetRootDir()
	//err := Cmd.Execute()
	//if err != nil {
	//	log.Fatal("启动错误")
	//}

	//Test.TestImplode()
}

func init() {
	fmt.Println("Main package init")
}

func test1() {
	now := time.Now()
	now.Second()

	aa := yaml.Decoder{}
	fmt.Println(aa)
	//fmt.Println(now.Year(), int(now.Month()), int(now.Day()), now.Hour(), now.Minute(), now.Second())
	//month := fmt.Sprintf("%02d", int(now.Month()))
	//fmt.Println(month)
	//
	//n := 88.88
	//fmt.Printf("%09.2f\n", n)
	//
	//fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	//fmt.Println("Formatted time:", now.Format("2006-01-02 15:04:05"))

	ticker := time.Tick(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker:
				fmt.Println("ticker time")

			}
		}
	}()

	ch := make(chan int, 1)
	go func(c chan int) {
		for {
			select {
			case ret := <-c:
				fmt.Println("接收成功", ret)
			}
		}
	}(ch) // 启用goroutine从通道接收值
	ch <- 10
	ch <- 11
	fmt.Println("发送成功")
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
		time.Sleep(1 * time.Hour)
		isErrDoneChan <- 1
	}()

	select {
	case <-isErrDoneChan:
		//发送消息给管理员
		fmt.Println("hhhhh")
	}
	time.Sleep(1 * time.Second)
}
