package System

import (
	"log"
	"os"
	"runtime"
	"strconv"
)

func IsDev() bool {
	return true
}

func IsTest() bool {
	return true
}

func IsProd() bool {
	return true
}

func IsWindows() bool {
	osName := runtime.GOOS
	switch osName {
	case "windows":
		return true
	default:
		return false
	}
}

func IsLinux() bool {
	osName := runtime.GOOS
	switch osName {
	case "linux":
		return true
	default:
		return false
	}
}

// IsMacos
//
//	@Description:
//	@return bool
func IsMacos() bool {
	osName := runtime.GOOS
	switch osName {
	case "darwin":
		return true
	default:
		return false
	}
}

// SaveMainPid 保存服务启动进程ID
func SaveMainPid() {
	pid := os.Getpid()
	serverFile, _ := os.Create(GetServerPidFile())
	_, err := serverFile.WriteString(strconv.Itoa(pid))
	if err != nil {
		log.Fatal("save server.pid error")
	}
}

// GetMainPid 获取主进程ID
func GetMainPid() int {
	pid, err := os.ReadFile(GetServerPidFile())
	if err != nil {
		log.Fatal("读取不到server.pid")
	}
	var serverPid int
	serverPid, _ = strconv.Atoi(string(pid))
	return serverPid
}
