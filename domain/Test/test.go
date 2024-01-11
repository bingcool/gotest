package Test

import (
	"fmt"
	"sync"
)

// TestFor 测试for循环体的变量
func TestFor() {
	var ids []*int
	for i := 0; i < 3; i++ {
		i = 10
		ids = append(ids, &i)
	}
	fmt.Println(ids)
	for _, id := range ids {
		fmt.Println(*id)
	}
}

// TestSyncMap 测试
func TestSyncMap() {
	var m sync.Map
	// 1. 写入
	m.Store("age1", 18)
	m.Store("age2", 20)

	// 2. 读取
	age, _ := m.Load("age1")

	//str := age + "年龄"
	//fmt.Println(str)
	myage, ok1 := age.(string)
	if ok1 {
		fmt.Println(myage)
	}

	// 3. 遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})

	// 4. 删除
	m.Delete("age1")
	age, ok := m.Load("age1")
	fmt.Println(age, ok)

	// 5. 读取或写入
	m.LoadOrStore("age2", 100)
	age, _ = m.Load("age2")
	fmt.Println(age)
}
