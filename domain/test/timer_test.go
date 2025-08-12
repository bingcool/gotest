package test

import (
	"fmt"
	"goTest/domain/system"
	"math"
	"runtime"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

func TestTick(t *testing.T) {
	// 获取时间戳
	nowTime := time.Now().Unix() + 11
	start_time := "2025-01-01"
	startDate, _ := time.Parse("2006-01-02", start_time)

	fmt.Println(startDate.Format("2006-01"))

	ticker1 := time.NewTicker(5 * time.Second)
	// 一定要调用Stop()，回收资源
	defer ticker1.Stop()

	// 循环中等待时间是否已达到
	go func(t *time.Ticker) {
		for {
			// 在该协程中阻塞等待，5s时间到了，就会触发一个chan返回，不影响其他协程的执行
			_ = <-t.C

			nowTime1 := time.Now().Unix()
			// 满足条件，人为控制stop定时器
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
			// 没有设置default分支，将在这里阻塞一直等待
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

	// 获取当前协程的数量
	coroutineNum := runtime.NumGoroutine()
	fmt.Println("当前协程数量：", coroutineNum)

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
		//default:
		//	fmt.Println("default")
	}
	time.Sleep(1 * time.Second)
}

func TestAfter(t *testing.T) {
	// 一次性定时任务
	afterTimer := time.AfterFunc(5*time.Second, func() {
		fmt.Println("after time")
	})

	defer afterTimer.Stop()

	time.Sleep(10 * time.Second)
}

func TestTicker(t *testing.T) {
	duration, _ := time.ParseDuration("1h33m29s")
	fmt.Println(duration.String())

	fmt.Println(duration.Seconds())

	fmt.Println(math.Round(duration.Minutes()*1000) / 1000)

	// 小数处理
	fmt.Println(math.Round(3.1415*1000) / 1000)

}

func TestTime1(tt *testing.T) {
	t := time.Now()
	fmt.Println("日期", t.Format("2006-01-02 15:04:05"))
	fmt.Println("时间戳：", t.Unix())

	y := t.Year()                 //年
	m := (int)(t.Month())         //月,转为整型
	d := t.Day()                  //日
	h := t.Hour()                 //小时
	i := t.Minute()               //分钟
	s := t.Second()               //秒
	fmt.Println(y, m, d, h, i, s) //2018 July 11 15 24 59
}

func TestTime2(tt *testing.T) {
	datetime := gtime.Date()
	fmt.Println(datetime)

	toTime, err := gtime.StrToTime("2023-07-05 15:04:05")
	if err != nil {
		return
	}

	fmt.Println(toTime.Timestamp())

	toTime = gtime.New()
	a1 := gtime.New(toTime.Timestamp()).Layout("2006-01-02 15:04:05")

	fmt.Println(a1)

	a2 := toTime.Weekday()

	fmt.Println(a2)
}

func TestTime3(tt *testing.T) {
	str := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(str)
	local, _ := time.LoadLocation("Asia/Shanghai")
	timeT, _ := time.ParseInLocation("2006-01-02 15:04:05", str, local)
	fmt.Println(timeT)
}
func TestTime4(tt *testing.T) {
	startTime, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	endTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(endTime.Format("2006-01-02 15:04:05"))
	fmt.Println(startTime.Format("2006-01-02 15:04:05"))
}
