package test

import (
	"fmt"
	"testing"
)

type CMAP[T int | float64] []T

// TestSelect3 泛型处理
func TestSelect3(t *testing.T) {
	// 声明一个有容量的切片，长度为1，容量为5，这时候0~长度的位置的值将初始化为零值
	cmap1 := make(CMAP[int], 1, 5)
	cmap1 = append(cmap1, 1)
	fmt.Println(cmap1)

	// 声明一个空的切片
	var cmap2 CMAP[float64]
	fmt.Println(cmap2)
	cmap2 = append(cmap2, 1.1)
	fmt.Println(cmap2)

	// 初始化切片并且初始化值
	cmap3 := CMAP[int]{1, 2, 3, 4}
	fmt.Println(cmap3)
}
