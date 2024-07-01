package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestP1(t *testing.T) {
	var a []any

	fmt.Println(a)
	fmt.Println(a...)
}

func TestArr1(t *testing.T) {
	x := [3]int{1, 2, 3}

	// 数组是值传递
	// 匿名函数, 传入数组, 尝试通过数组索引修改数组
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println("arr:", arr)
	}(x)

	fmt.Println("x:", x)
}

func TestGo(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	//for {
	//	// 调度函数
	//	runtime.Gosched()
	//}

	// 或者使用select
	select {
	//default:
	}
}

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

// 测试defer,发生参数逃逸
func TestGo1(t *testing.T) {
	n := 5
	f := func() int {
		// n发生参数逃逸
		n++
		return n
	}

	// defer是先入后出，最后执行这个defer
	defer func() {
		fmt.Println("go1")
		fmt.Println(n)
		fmt.Println(f())
	}()

	// 先执行这个defer
	defer fmt.Println(f())
	i := f()
	fmt.Println(i)
}
