package cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"goTest/domain/crontab"
	"goTest/domain/daemon"
	"goTest/domain/system"
	"os"
	"os/exec"
)

// go run main.go daemon-start consume-user-order --name=bingcool

var daemonStartCommandName = "daemon-start"

var DaemonStartCmd = &cobra.Command{
	Use:   daemonStartCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		console.NewConsole().PutCommand(cmd)
		startDaemon(cmd, args)
	},
}

func init() {
	initDaemonStartFlags()
}

func initDaemonStartFlags() {
	if os.Args[1] == daemonStartCommandName {
		if len(os.Args) < 3 {
			panic("请指定启动进程名")
		}

		processName := os.Args[2]
		scheduleList := getSchedule()
		processItemMap, isExist := scheduleList[processName]
		if !isExist {
			panic("找不到对应的进程名=" + processName)
		}
		flagsFn, isExistFlagFn := processItemMap["flags"]
		if !isExistFlagFn {
			panic(processName + "找不到对应的flags")
		}

		// cli 传进来的flag参数优于配置文件的自定义的参数
		if len(os.Args) > 2 {
			parseFlag(DaemonStartCmd, os.Args[2:])
		}

		// 配置文件flag参数
		flags := flagsFn(DaemonStartCmd)
		parseFlag(DaemonStartCmd, flags)

		initDaemonFlags(DaemonStartCmd)
	}
}

func getSchedule() console.ScheduleType {
	var scheduleList console.ScheduleType
	if isFromCron() {
		scheduleList = *crontab.RegisterCronSchedule()
	} else {
		scheduleList = *daemon.RegisterDaemonSchedule()
	}

	return scheduleList
}

func startDaemon(cmd *cobra.Command, args []string) {
	processName := args[0]
	scheduleList := getSchedule()
	processItemMap, isExist := scheduleList[processName]
	if !isExist {
		panic("找不到对应的进程名=" + processName)
	}

	fn, isExistFn := processItemMap["fn"]
	if !isExistFn {
		panic(processName + "找不到对应的处理函数")
	}

	if isFork(cmd) {
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

	if isFromCron() {
		cronHandle(processName, fn, cmd)
	} else {
		daemonHandle(processName, fn, cmd)
	}
}

func daemonHandle(processName string, fn func(cmd *cobra.Command) []string, cmd *cobra.Command) {
	createDaemonPidPath()
	saveProcessPid(getDaemonPidFile(processName))
	channel := make(chan int, 1)
	go func(channel chan int) {
		fn(cmd)
	}(channel)

	cronSchedule := cron.New()
	_, _ = cronSchedule.AddFunc("@every 2s", func() {
		saveProcessPid(getDaemonPidFile(processName))
	})
	cronSchedule.Start()

	select {
	case <-channel:
	}
}

func forkDaemonProcess(args []string) {
	if system.IsLinux() || system.IsMacos() {
		processName := args[0]
		newArgs := make([]string, 0)
		newArgs = append(newArgs, daemonStartCommandName, processName)
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
