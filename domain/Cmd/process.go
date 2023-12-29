package Cmd

import (
	"goTest/domain/System"
	"goTest/domain/Util"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

func saveProcessPid(pidFile string) {
	pid := os.Getpid()
	serverFile, _ := os.Create(pidFile)
	_, err := serverFile.WriteString(strconv.Itoa(pid))
	if err != nil {
		log.Fatal("save server.pid error")
	}
}

func getDaemonPidFile(processName string) string {
	pidFile := filepath.Join(getDaemonPidPath(), processName+".pid")
	return pidFile
}

func getCronPidFile(processName string) string {
	pidFile := filepath.Join(getCronPidPath(), processName+".pid")
	return pidFile
}

func getProcessPid(pidFilePath string) int {
	pid, err := os.ReadFile(pidFilePath)
	if err != nil {
		return 0
	}
	var serverPid int
	serverPid, _ = strconv.Atoi(string(pid))
	return serverPid
}

func getDaemonPidPath() string {
	folderPath := filepath.Join(System.GetDaemonPath(), "pid")
	return folderPath
}

func getCronPidPath() string {
	folderPath := filepath.Join(System.GetCronPath(), "pid")
	return folderPath
}

func createPidPath(pidPath string) {
	dirPath := filepath.Join(pidPath, "pid")
	// 使用 Stat 函数获取文件夹信息
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Println("创建Pid文件夹失败")
		}
	}
}

func createDaemonPidPath() {
	createPidPath(System.GetDaemonPath())
}

func createCronPidPath() {
	createPidPath(System.GetCronPath())
}

func isFork(args []string) bool {
	if Util.ContainsInSlice(args, "d") || Util.ContainsInSlice(args, "D") {
		return true
	}
	return false
}

func isProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		log.Printf("Process Id=%d is dead!", pid)
		return false
	} else {
		log.Printf("Process Id=%d is alive!", pid)
		return true
	}
}
