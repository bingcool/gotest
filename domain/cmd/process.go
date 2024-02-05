package cmd

import (
	"flag"
	"github.com/spf13/cobra"
	"goTest/domain/system"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

var cronServerPidFile = "crontab-server.pid"

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
	pidPath := filepath.Join(system.GetDaemonPath(), "pid")
	_, err := os.Stat(pidPath)
	if err != nil {
		createPidPath(pidPath)
	}
	return pidPath
}

func getCronPidPath() string {
	pidPath := filepath.Join(system.GetCronPath(), "pid")
	_, err := os.Stat(pidPath)
	if err != nil {
		createPidPath(pidPath)
	}
	return pidPath
}

func saveCronServerPid(pid int) {
	cronServerPidFile := filepath.Join(getCronPidPath(), cronServerPidFile)
	serverFile, _ := os.Create(cronServerPidFile)
	_, err := serverFile.WriteString(strconv.Itoa(pid))
	if err != nil {
		log.Fatal("save crontab-server.pid error")
	}
}

func getCronServerPid() int {
	cronServerPidFile := filepath.Join(getCronPidPath(), cronServerPidFile)
	pid, err := os.ReadFile(cronServerPidFile)
	if err != nil {
		log.Fatal("读取不到cron-server.pid")
	}
	var serverPid int
	serverPid, _ = strconv.Atoi(string(pid))
	return serverPid
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
	createPidPath(system.GetDaemonPath())
}

func createCronPidPath() {
	createPidPath(system.GetCronPath())
}

func isFork(cmd *cobra.Command) bool {
	if value, _ := cmd.Flags().GetInt("daemon"); value == 1 {
		return true
	}
	return false
}

var daemonFlag int

func initDaemonFlags(cmd *cobra.Command) {
	// 初始化系统flags的默认值
	if cmd.Flags().Lookup("daemon") == nil {
		cmd.Flags().IntVar(&daemonFlag, "daemon", 0, "--fork-cron=0 or --fork-cron=1")
	}
}

func isProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	} else {
		return true
	}
}

func isFromCron() bool {
	flag.Parse()
	list := flag.Args()
	for _, item := range list {
		if item == "--from-flag=cron" {
			return true
		}
	}
	return false
}
