package test

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("kkkkkkkkkk")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("panic")
}
