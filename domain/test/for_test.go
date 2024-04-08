package test

import (
	"fmt"
	"testing"
	"time"
)

func TestFor(t *testing.T) {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04::05"))
	}
}

func TestForSelect(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("receive")
				time.Sleep(5 * time.Second)
				fmt.Println(time.Now().Format("2006-01-02 15:04::05"))
			}
		}
	}()

	time.AfterFunc(10*time.Second, func() {
		ch <- 1
	})
	<-make(chan struct{})

}
