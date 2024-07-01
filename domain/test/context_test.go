package test

import (
	"fmt"
	"goTest/domain/system"
	"golang.org/x/net/context"
	"os"
	"runtime/debug"
	"syscall"
	"testing"
	"time"
)

func TestCtx1(t *testing.T) {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, "requestID", "12345")
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			}
		}

	}()

	time.Sleep(time.Second * 3)

	cancel()

	time.Sleep(time.Second * 10)

}

type myint int

func TestCtx2(t *testing.T) {
	go func() {
		fmt.Println("协程处理完毕")
	}()

	fmt.Println("进程ID", os.Getpid())
	system.Signal(syscall.SIGTERM, func(sigs os.Signal) {
		fmt.Println("收到信号", sigs)
	})

	v := myint(3)
	fmt.Println(v)

	stack := debug.Stack()
	fmt.Println(string(stack))
	time.Sleep(30 * time.Second)

}
