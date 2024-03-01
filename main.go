package main

import (
	"fmt"
	"goTest/domain/cmd"
	"goTest/domain/system"
	"gopkg.in/yaml.v2"
	"time"
)

func main() {

	system.GetRootDir()
	_ = cmd.Execute()

	//test.TestImplode()
}

func init() {
	//fmt.Println("Main package init")
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
