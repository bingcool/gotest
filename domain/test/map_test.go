package test

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/samber/lo"
	"log"
	"sync"
	"testing"
)

func TestMapSet(t *testing.T) {
	//var arr1 map[string]string
	arr1 := make(map[string]string)
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
	// 数组初始化将自动生复制对应类型的零值
	var a [5]int
	fmt.Println(a)
}

func TestContains(t *testing.T) {
	has := lo.Contains([]int{1, 2, 3, 4, 5}, 30)
	fmt.Println(has)

	// 自定义函数实现判断
	has1 := lo.ContainsBy([]int{1, 2, 30, 4, 5}, func(item int) bool {
		if item == 30 {
			return true
		}
		return false
	})

	fmt.Println(has1)
}

// 并发map发生异常
// fatal error: concurrent map writes
func TestParallelMap(t *testing.T) {
	m := make(map[int]string)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer func() {
				if e := recover(); e != nil {
					log.Printf("recover: %v", e)
				}
			}()

			m[i] = "Go 语言编程之旅：一起用 Go 做项目"
		}(i)
	}
}

// TestSyncMap 测试
func TestSyncMap(t *testing.T) {

	var m sync.Map
	// 1. 写入
	m.Store("age1", 18)
	m.Store("age2", 20)

	// 2. 读取
	age, _ := m.Load("age1")

	// any的类型可以通过fmt.Sprintf转实际类型
	str := fmt.Sprintf("%v", age) + "年龄"
	fmt.Println(str)
	//myage, ok1 := age.(string)

	// 3. 遍历(依然是无序的)
	m.Range(func(key, value any) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})

	// 4. 删除
	m.Delete("age1")
	age1, ok := m.Load("age1")
	fmt.Println(age1, ok)

	// 5. 读取或写入
	m.LoadOrStore("age2", 100)
	age2, _ := m.Load("age2")
	fmt.Println(age2)

}
