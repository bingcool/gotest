package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"goTest/domain/Middlewares"
	"goTest/domain/Router"
	"log"
	"runtime/pprof"
)

var threadProfile = pprof.Lookup("threadcreate")

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "My Gin application",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func startServer() {
	//runtime.GOMAXPROCS(2)
	fmt.Printf(("threads in starting Num: %d\n"), threadProfile.Count())
	//for {
	//	time.Sleep(1 * time.Second) // 休眠1秒钟
	//}

	//fmt.Println(Middlewares.GetCurrentGoroutineID())

	r := gin.Default()
	// 设置日志中间件，主要用于打印请求日志
	// 设置Recovery中间件，主要用于拦截panic错误，不至于导致进程崩掉

	defer func() {
		fmt.Println("start defer")
	}()

	Middlewares.SetGlobalMiddleware(r)

	Router.SetupProductRouter(r)
	Router.SetOrderRouter(r)

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
