package test

import (
	"fmt"
	"github.com/juju/errors"
	"github.com/shirou/gopsutil/v3/mem"
	"goTest/domain/system"
	"runtime"
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

func TestTick(t *testing.T) {
	nowTime := time.Now().Unix() + 11

	ticker1 := time.NewTicker(5 * time.Second)
	// 一定要调用Stop()，回收资源
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			// 在改协程中阻塞等待，不影响其他协程的执行
			_ = <-t.C

			nowTime1 := time.Now().Unix()
			// 满足条件，认为控制stop定时器
			if nowTime1 > nowTime {
				t.Stop()
			} else {
				fmt.Println("Ticker:", time.Now().Format("2006-01-02 15:04:05"))
			}

		}
	}(ticker1)

	println("Ticker start")
	time.Sleep(30 * time.Second)
	fmt.Println("ok")
}

func TestTime(t *testing.T) {
	ticker := time.Tick(1 * time.Second)
	// 在协程中创建一个定时器，不影响往下面的流程的执行
	go func() {
		for {
			select {
			case <-ticker:
				fmt.Println("ticker time")
			default:

			}
		}
	}()

	// 创建channel,协程间通信
	ch := make(chan int, 1)
	go func(c chan int) {
		for {
			select {
			case ret := <-c:
				fmt.Println("接收成功", ret)
			}
		}
	}(ch) // 启用goroutine从通道接收值
	ch <- 10
	ch <- 11

	time.Sleep(5 * time.Second)
}

func TestSignal(t *testing.T) {
	numCPU := runtime.NumCPU()
	fmt.Println("Number of CPUs:", numCPU)

	// 设置线程数量为 4
	runtime.GOMAXPROCS(1)

	// 信号监听
	system.EventLoopSigtermSignal()

	fmt.Println("start start ")
	//debug.SetMaxThreads(10)

	isErrDoneChan := make(chan int, 1)
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("开始向channel发送消息")
		isErrDoneChan <- 1
	}()

	select {
	// 没设置default,则一直阻塞等待。设置了，如果channel没有数据就绪,直接执行default。同时子协程也会退出
	case <-isErrDoneChan:
		//发送消息给管理员
		fmt.Println("channel接受到信息")
	default:
		fmt.Println("default")
	}
	time.Sleep(1 * time.Second)
}
