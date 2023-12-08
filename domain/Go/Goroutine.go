package Go

import "fmt"

type goFunc func(ctx Context)

type Context struct {
	BucketMap map[string]any
	Channel   chan any
}

func Run(callback goFunc, ctx Context) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered:", r)
			}
		}()
		callback(ctx)
		fmt.Println("this is goApp")
	}()
}
