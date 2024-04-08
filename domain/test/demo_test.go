package test

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/shirou/gopsutil/v3/mem"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestFor 测试for循环体的变量
func TestForLoop(t *testing.T) {
	var ids []int
	for i := 0; i < 3; i++ {
		ids = append(ids, i)
	}
	fmt.Println(ids)
	for _, id := range ids {
		fmt.Println(id)
	}
}

func TestError(t *testing.T) {
	errors.New("test error")
	fmt.Println(fmt.Errorf("error test"))
}

// TestSyncMap 测试
func TestSyncMap(t *testing.T) {
	var m sync.Map
	// 1. 写入
	m.Store("age1", 18)
	m.Store("age2", 20)

	// 2. 读取
	age, _ := m.Load("age1")

	//str := age + "年龄"
	//fmt.Println(str)
	//myage, ok1 := age.(string)

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

// TestGopsutil 内存使用统计
func TestGopsutil(t *testing.T) {
	v, _ := mem.VirtualMemory()
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}

var counter int
var mutex sync.Mutex

// TestMurexLock 测试互斥锁 哪个协程获取到锁哪个协程就执行
func TestMurexLock(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 300; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
	fmt.Println("Counter:", counter)
}

func increment(wg *sync.WaitGroup) {
	// 多个协程并发，counter累加必须加互斥锁。此处获取锁
	mutex.Lock()
	counter++
	defer func() {
		// 此处释放锁
		mutex.Unlock()
		wg.Done()
	}()
}

var value int32

func TestAutomic(t *testing.T) {
	// 原子写入操作
	atomic.StoreInt32(&value, 0)
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt32(&value, 1)
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(value)

}

type OrderModel struct {
	OrderId uint32
}

// TestStruct 测试结构体初始化
func TestStruct(t *testing.T) {
	//字面量初始化
	order := OrderModel{}
	//零值初始化
	var order1 OrderModel
	fmt.Println(order)
	fmt.Println(order1)
}

// TestStringIndex 字符串分割处理
func TestStringIndex(t *testing.T) {
	username := "bingcool@email.com"
	// 获取字符所在字符串的索引位置，并返回
	idx := strings.Index(username, "@")
	var name string
	// 存在这个字符并且返回，那么这个idx将不会为-1
	if idx != -1 {
		// 读取这个字符串的切片数组
		name = username[:idx]
	} else {
		name = username
	}

	fmt.Printf("获取分割的前部分字符:%s\n", name)

	// 使用1.18版本后新增的方法strings.Cut()
	before, after, ok := strings.Cut(username, "@")
	if !ok {
		panic("分割报错了")
	}

	fmt.Println(before)
	fmt.Println(after)
}

// TestImplode 字符串与字符串切片之间的转换
func TestImplode(t *testing.T) {
	// 切片合并成字符串
	items := []string{"中国", "人民", "解放军"}
	str := strings.Join(items, "@")
	fmt.Println(str)

	// 字符串分割成切片
	items = strings.Split(str, "@")
	for _, v := range items {
		fmt.Println(v)
	}
}
