package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"runtime/pprof"
)

var threadProfile = pprof.Lookup("threadcreate")

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Printf(("threads in starting: %d\n"), threadProfile.Count())
	//for {
	//	time.Sleep(1 * time.Second) // 休眠1秒钟
	//}

	r := gin.Default()
	// 设置日志中间件，主要用于打印请求日志

	r.Use(gin.Logger())
	// 设置Recovery中间件，主要用于拦截panic错误，不至于导致进程崩掉
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
