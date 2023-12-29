package System

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HandleSignal func(sigs os.Signal)

func Signal(sigs os.Signal, fn HandleSignal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sigs)
	go func() {
		defer close(ch)
		for {
			s := <-ch
			if s == sigs {
				fn(sigs)
			}
		}
	}()
}

// EventLoopSigtermSignal 信号15平滑kill进程
func EventLoopSigtermSignal() {
	Signal(syscall.SIGTERM, func(sigs os.Signal) {
		fmt.Println("开始推出进程Sigterm")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	})

	Signal(syscall.SIGINT, func(sigs os.Signal) {
		fmt.Println("开始推出进程Sigint")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	})
}
