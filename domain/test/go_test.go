package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func TestGo2(t *testing.T) {
	// 通过设置不同的逻辑处理核数，可以观察到 go 协程的并发执行
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
