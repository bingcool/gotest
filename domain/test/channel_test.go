package test

import (
	"fmt"
	"testing"
	"time"
)

// TestChannel 无缓存的channel
func TestChannel(t *testing.T) {
	ch := make(chan int)

	//go func() {
	//	for {
	//		// 阻塞
	//		v := <-ch
	//		t.Log("接收到数据", v)
	//	}
	//}()

	// 如果没有消费，将会一直阻塞
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func TestChannel2(t *testing.T) {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		t.Log("push-", i)
		ch <- i
	}

	go func() {
		for {
			select {
			case v := <-ch:
				t.Log("接收到数据", v)
			}
		}
	}()

	// 阻塞
	time.Sleep(5 * time.Second)
}

func TestChannel3(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 5   // 写入一个值
	close(ch) // 关闭通道

	value, ok := <-ch      // 读取值，ok为true，value为5
	fmt.Println(value, ok) // 输出: 5 true

	// 通道关闭后，通道内没有数据，读取时，ok为false，value为0（int类型的零值）
	value, ok = <-ch       // 再次读取，ok为false，value为0（int类型的零值）
	fmt.Println(value, ok) // 输出: 0 false
}
