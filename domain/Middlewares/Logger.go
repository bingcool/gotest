package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
)

func SetLoggerWithWriter(router *gin.Engine) {
	currentDir, _ := os.Getwd()
	//currentDir = filepath.ToSlash(currentDir)
	logFilePath := filepath.Join(currentDir, "domain", "Storage", "gin.log")
	// 如果你想将日志输出到文件中，可以将输出重定向到一个文件
	logFile, err := os.Create(logFilePath)
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
