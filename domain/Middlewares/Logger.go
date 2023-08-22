package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func SetLoggerWithWriter(router *gin.Engine) {
	currentDir, _ := os.Getwd()
	//currentDir = filepath.ToSlash(currentDir)
	logFilePath := filepath.Join(currentDir, "domain", "Storage", "gin.log")
	// 如果你想将日志输出到文件中，可以将输出重定向到一个文件
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func(logFile *os.File) {

	}(logFile)

	//fmt.Println(GetCurrentGoroutineID())
	// 设置日志输出到文件
	gin.DefaultWriter = io.MultiWriter(logFile)

	// 设置请求时触发的中间件
	router.Use(LoggerToFile(logFile))
}

func GetCurrentGoroutineID() uint64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to get goroutine ID:", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idStr := string(buf[:n])

	var id uint64
	_, err := fmt.Sscanf(idStr, "goroutine %d", &id)
	if err != nil {
		fmt.Println("Failed to parse goroutine ID:", err)
		return 0
	}

	return id
}

// LoggerToFile 自定义日志中间件
func LoggerToFile(logFile *os.File) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求信息
		_, err := logFile.WriteString("Request: " + c.Request.URL.Path + "\n")
		if err != nil {
			return
		}

		// 继续处理请求
		c.Next()

		// 记录响应信息
		_, err = logFile.WriteString("Response: " + "kkkkkkkk" + "\n")
		if err != nil {
			return
		}
	}
}
