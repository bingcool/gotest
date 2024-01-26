package Go

import (
	"fmt"
	"runtime"
)

type goFunc func(ctx Context)

type Context struct {
	BucketMap map[string]any
	Channel   chan any
}

func Run(callback goFunc, ctx Context) {
	go func() {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
		callback(ctx)
		fmt.Println("this is goApp")
	}()
}

func GetGoroutineID() uint64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to get goroutine ID:", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idStr := string(buf[:n])

	var id uint64
	_, err := fmt.Sscanf(idStr, "goroutine %d", &id)
	if err != nil {
		fmt.Println("Failed to parse goroutine ID:", err)
		return 0
	}

	return id
}
