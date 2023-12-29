package Cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Daemon"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var daemonStartCommandName = "daemon-start"

var DaemonStartCmd = &cobra.Command{
	Use:   daemonStartCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		// log.Printf("daemon before run ")
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		startDaemon(cmd, args)
	},
}

func init() {
	initDaemonStartFlags(DaemonStartCmd)
	initParseFlag(daemonStartCommandName, DaemonStartCmd)
}

func initDaemonStartFlags(cmd *cobra.Command) {
	if os.Args[1] == daemonStartCommandName {
		processName := os.Args[2]
		scheduleList := *Daemon.RegisterDaemonSchedule()
		processItemMap, isExist := scheduleList[processName]
		if !isExist {
			panic("找不到对应的进程名=" + processName)
		}

		flagsFn, isExistFlagFn := processItemMap["flags"]
		if !isExistFlagFn {
			panic(processName + "找不到对应的flags")
		}
		flags := flagsFn(cmd)
		log.Println(flags)
		parseFlag(cmd, flags)
	}
}

func startDaemon(cmd *cobra.Command, args []string) {
	processName := args[0]
	scheduleList := *Daemon.RegisterDaemonSchedule()
	processItemMap, isExist := scheduleList[processName]
	if !isExist {
		panic("找不到对应的进程名=" + processName)
	}

	fn, isExistFn := processItemMap["fn"]
	if !isExistFn {
		panic(processName + "找不到对应的处理函数")
	}

	if isFork(args) {
		forkDaemonProcess(args)
		os.Exit(0)
	} else {
		startProcess(processName, fn, cmd)
	}
}

func startProcess(processName string, fn func(cmd *cobra.Command) []string, cmd *cobra.Command) {
	// 判断进程是否已经启动了
	//pid := getProcessPid(getDaemonPidFile(processName))
	//if pid > 0 {
	//	if isProcessRunning(pid) {
	//		log.Printf("进程ID=%d已经启动，无需重新启动", pid)
	//		return
	//	}
	//}
	createDaemonPidPath()
	saveProcessPid(getDaemonPidFile(processName))

	channel := make(chan int, 1)
	go func(channel chan int) {
		fn(cmd)
	}(channel)

	c := cron.New()
	_, _ = c.AddFunc("@every 2s", func() {
		saveProcessPid(getDaemonPidFile(processName))
	})
	c.Start()

	select {
	case <-channel:
	}
}

func forkDaemonProcess(args []string) {
	osName := runtime.GOOS
	switch osName {
	// linux，macos
	case "linux", "darwin":
		newArgs := make([]string, 0)
		newArgs = append(newArgs, daemonStartCommandName)
		for _, value := range args {
			if value != "d" && value != "D" {
				newArgs = append(newArgs, value)
			}
		}
		newCmd := exec.Command(os.Args[0], newArgs...)
		newCmd.Stdin = os.Stdin
		//newCmd.Stdout = os.Stdout
		newCmd.Stderr = os.Stderr
		err := newCmd.Start()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "创建exec守护进程失败: %s\n", err)
		}
	}
}
