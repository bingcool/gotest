package Cmd

import (
	"goTest/domain/System"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

func saveDaemonProcessPid(commandName string) {
	pidFile := getPidFilePath(commandName)
	pid := os.Getpid()
	serverFile, _ := os.Create(pidFile)
	_, err := serverFile.WriteString(strconv.Itoa(pid))
	if err != nil {
		log.Fatal("save server.pid error")
	}
}

func getPidFilePath(commandName string) string {
	pidFile := filepath.Join(getDaemonPidPath(), commandName+".pid")
	return pidFile
}

func getDaemonProcessPid(pidFilePath string) int {
	pid, err := os.ReadFile(pidFilePath)
	if err != nil {
		return 0
	}
	var serverPid int
	serverPid, _ = strconv.Atoi(string(pid))
	return serverPid
}

func getDaemonPidPath() string {
	folderPath := filepath.Join(System.GetDaemonPath(), "Pid")
	return folderPath
}

func createDaemonPidPath() {
	folderPath := filepath.Join(System.GetDaemonPath(), "Pid")
	// 使用 Stat 函数获取文件夹信息
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			log.Println("创建Pid文件夹失败")
		}
	}
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
