package test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestJson(t *testing.T) {
	str := `{name:'bingcool'}`
	strByte := []byte(str)
	var a any
	err := json.Unmarshal(strByte, &a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(a)
}

func TestJson2(t *testing.T) {
	var a string
	b := 1
	a = fmt.Sprintf("%v", b)

	fmt.Println(a)

}

func TestJson3(t *testing.T) {
	var intValue = 10
	var floatValue = 3.14

	// 可以判断是否是数字
	c, err := strconv.ParseFloat("-0.3", 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if reflect.TypeOf(c).Kind() == reflect.Float64 {
		fmt.Println("c---是整数")
	}

	fmt.Printf("intValue 的类型是：%s\n", reflect.TypeOf(intValue))
	if reflect.TypeOf(intValue).Kind() == reflect.Int {
		fmt.Println("intValue 是整数")
	}

	fmt.Printf("floatValue 的类型是：%s\n", reflect.TypeOf(floatValue))
	if reflect.TypeOf(floatValue).Kind() == reflect.Float64 {
		fmt.Println("floatValue 是浮点数")
	}
}
