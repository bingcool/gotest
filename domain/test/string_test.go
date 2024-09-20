package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestString1(t *testing.T) {
	s := "huang zengbing黄"
	fmt.Println(s)

	// 读取字符串的长度，返回的是字节大小
	fmt.Println(len(s))

	// 读取字符串的长度，返回的是字符个数
	runeString := []rune(s)
	fmt.Println(len(runeString))

	// 打印每个字符
	for i := 0; i < len(runeString); i++ {
		fmt.Printf("%c", runeString[i])
	}

	// 字符串拼接用+
	s1 := s + "bbb"
	fmt.Println(s1)

	s2 := fmt.Sprintf("%s%s", s, "bbb")

	fmt.Println(s2)

	// bytes.Buffer 提供了更高效的字符串拼接方法，尤其是当拼接操作发生在循环中时
	b := &strings.Builder{}
	b.WriteString("Hello, ")
	b.WriteString("world!")
	s3 := b.String()
	fmt.Println(s3) // 输出: Hello, world!

	s4 := strings.Join([]string{"a", "b", "c"}, "")
	fmt.Println(s4)

}

func TestString2(t *testing.T) {
	s := "huang zeng bing"
	fmt.Println(strings.Contains(s, "zeng"))

	fmt.Println(strings.HasPrefix(s, "hua"))

	s1 := strings.Replace(s, "zeng", "zzzzzzzz", 1)
	fmt.Println(s1)

	// 返回子串在字符串中首次出现的位置，从0开始
	fmt.Println(strings.Index(s, "h"))
	// 返回子串在字符串中最后一次出现的位置
	fmt.Println(strings.LastIndex(s, "g"))

	fmt.Println(strings.ToUpper(s))

}
