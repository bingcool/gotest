package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"log"
	"os"
	"syscall"
)

var daemonStopCommandName = "daemon-stop"

var DaemonStopCmd = &cobra.Command{
	Use:   daemonStopCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		// log.Printf("daemon before run ")
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		stopDaemon(args)
	},
}

func init() {
	initParseFlag(daemonStopCommandName, DaemonStopCmd)
}

func stopDaemon(args []string) {
	processName := args[0]
	pidFile := getDaemonPidFile(processName)
	pid := getProcessPid(pidFile)
	killProcess(pid, pidFile)
}

func killProcess(pid int, filePath string) {
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Printf("Error finding process:", err)
		if len(filePath) > 0 {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			log.Printf("File deleted successfully")
		}
	} else {
		if isProcessRunning(pid) {
			_ = process.Signal(syscall.SIGTERM)
			log.Printf("process stop successfully, file=%s", filePath)
		} else {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			log.Printf("进程不存在，删除失效的pid文件=%s", filePath)
		}
	}
}
