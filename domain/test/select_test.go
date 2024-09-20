package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	// 设置容量为1
	ch := make(chan int, 1)

	go func() {
		fmt.Println("等待5s")
		time.Sleep(5 * time.Second)
		ch <- 1
	}()

	defer close(ch)

	select {
	case v := <-ch:
		fmt.Println(" 5s 后 received data", v)
		break
		//default:
		//	fmt.Println("no value received")
	}

	// time.Sleep(2 * time.Second)

}

func TestSelect2(t *testing.T) {
	// 定义一个没有容量的通道
	ch := make(chan struct{})
	fmt.Println("go go go")
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}()
	fmt.Println("start push")
	// 因没有设置容量，下面向通道赋值，此通道将一直阻塞进程，不会退出
	ch <- struct{}{}
}
