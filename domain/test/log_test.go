package test

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"testing"
	"time"
)

func TestLoG(t *testing.T) {
	// 创建一个Lumberjack实例来配置日志分割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   generateLogFilename(), // 日志文件路径
		MaxSize:    100,                   // 单个日志文件的最大大小（以MB为单位）
		MaxBackups: 3,                     // 保留的旧日志文件的最大个数
		MaxAge:     7,                     // 保留的旧日志文件的最大天数
		Compress:   true,                  // 是否启用日志文件的压缩
	}

	// 创建一个Zap核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // JSON编码器
		zapcore.AddSync(lumberjackLogger),                        // 将日志写入Lumberjack实例
		zap.DebugLevel,                                           // 日志级别
	)

	// 创建一个Zap日志记录器
	logger := zap.New(core)

	// 示例日志记录
	logger.Info("This is an example log message.",
		zap.String("name", "bingcool"),
		zap.Int("age", 18),
	)

	// 关闭日志记录器
	err := logger.Sync()
	if err != nil {
		return
	}
}

func generateLogFilename() string {
	currentTime := time.Now()
	// 按照需要的命名格式生成日志文件名，例如：log_2022-01-01_12-30-45.log
	filename := fmt.Sprintf("log_%s.log", currentTime.Format("20060102"))
	return filename
}
