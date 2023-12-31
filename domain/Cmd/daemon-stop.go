package Cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"log"
	"os"
	"syscall"
)

// go run main.go daemon-stop consume-user-order

var daemonStopCommandName = "daemon-stop"

var DaemonStopCmd = &cobra.Command{
	Use:   daemonStopCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("请指定停止进程名")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		stopDaemon(args)
	},
}

func init() {
	initDaemonStopFlag()
}

func initDaemonStopFlag() {
	if os.Args[1] == daemonStopCommandName {
		if len(os.Args) > 3 {
			parseFlag(DaemonStopCmd, flag.Args()[3:])
		}
	}
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
