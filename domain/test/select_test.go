package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch := make(chan int, 1)

	go func() {
		ch <- 1
	}()

	defer close(ch)

	select {
	case v := <-ch:
		fmt.Println("received", v)
		break
	}

	time.Sleep(2 * time.Second)

}

func TestSelect2(t *testing.T) {
	// 定义一个没有容量的通道
	ch := make(chan struct{})
	fmt.Println("ggg")
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}()
	// 因没有容量，此通道将一直阻塞进程，不会退出
	ch <- struct{}{}
}
