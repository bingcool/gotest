package test

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
)

func TestMapSet(t *testing.T) {
	//var arr1 map[string]string
	arr1 := map[string]string{}
	fmt.Println(arr1)

	arr1["name"] = "bingcool"
	arr1["age"] = "33"

	referMap(arr1)

	// arr1作为参数传给某个函数后，可以发现age发生了改变，所以map的传参时引用传参
	fmt.Println(arr1)
}

func referMap(arr map[string]string) {
	arr["address"] = "深圳"
	arr["age"] = "35"

	fmt.Println(arr)
}

func GenerateMap() map[string]string {
	//var arr1 map[string]string
	arr1 := map[string]string{}
	arr1["name"] = "bingcool"
	arr1["age"] = "33"
	arr1["address"] = "深圳"
	arr1["sex"] = "1"
	arr1["pointer"] = "only"

	return arr1
}

func TestMapForeach(t *testing.T) {
	arr1 := GenerateMap()
	// map的循环输出是无序的，因为go底层是随机选择一个bucket作为开始循环的起始位置的
	for i, v := range arr1 {
		fmt.Println(i, ":", v)
	}
}

func TestMapCovertToJson(t *testing.T) {
	arr := GenerateMap()
	fmt.Println(arr)

	// map转json
	json := gconv.String(arr)
	fmt.Println(json)

	type User struct {
		Name    string
		Age     string
		Address string
		Sex     string
		Pointer string
	}

	var user *User
	// map转struct
	_ = gconv.Struct(arr, &user)
	g.Dump(user)

	var user1 *User
	// json字符传传struct对象
	_ = gconv.Struct(json, &user1)
	g.Dump(user1)

	//var map1 map[string]string
	map1 := map[string]string{}
	// json字符传传map
	_ = gconv.Struct(json, &map1)

	fmt.Println(map1["age"])

}

func TestMapToStruct(t *testing.T) {
	type User struct {
		Uid            int    `c:"uid"` // 这个tag代表是转为map之后的field
		Name           string `c:"name"`
		AddressCountry string `c:"address_country"`
	}

	user := &User{
		Uid:  1,
		Name: "john",
	}

	// struct对象转map
	userMap1 := gconv.Map(user)
	g.Dump(userMap1["uid1"])

	// struct对象指针转map
	g.Dump(gconv.Map(&User{
		Uid:  1,
		Name: "john",
	}))

	// 任意map类型
	g.Dump(gconv.Map(map[int]int{
		100: 10000,
	}))
}

func TestMapsToStruct(t *testing.T) {
	type User struct {
		Uid            int    `c:"uid"` // 这个tag代表是转为map之后的field
		Name           string `c:"name"`
		AddressCountry string `c:"address_country"`
	}

	user := &User{
		Uid:  1,
		Name: "john",
	}

	mapArr := gconv.Maps(user)

	fmt.Println(mapArr)

}

func TestArr(t *testing.T) {
	// 数组初始化对应类型的零值
	var a [5]int
	fmt.Println(a)
}
