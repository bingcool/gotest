package System

import (
	"path/filepath"
	"strings"
)

const rootDirLinux = "/home/wwwroot/go/goTest"
const rootDirMacos = "/Users/huangzengbing/Documents/wwwroot/goTest"
const rootDirWindows = "D:\\wwwroot\\go\\goTest"

func GetRootDir() string {
	items := make([]string, 0)
	separator := string(filepath.Separator)
	var rootPath = ""
	if IsWindows() {
		rootPath = rootDirWindows
	} else if IsLinux() {
		items = strings.Split(rootDirLinux, separator)
		rootPath = separator + filepath.Join(items...)
	} else {
		items = strings.Split(rootDirMacos, separator)
		rootPath = separator + filepath.Join(items...)
	}

	return rootPath
}

func GetStoragePath() string {
	storagePath := filepath.Join(GetRootDir(), "domain", "Storage")
	return storagePath
}

func GetStorageGinLog() string {
	logFilePath := filepath.Join(GetStoragePath(), "gin.log")
	return logFilePath
}

func GetServerPidFile() string {
	pidFile := filepath.Join(GetRootDir(), "server.pid")
	return pidFile
}

func GetDaemonPath() string {
	daemonPath := filepath.Join(GetRootDir(), "domain", "Daemon")
	return daemonPath
}

func MyTestAdd(a int, b int) int {
	return a + b
}
