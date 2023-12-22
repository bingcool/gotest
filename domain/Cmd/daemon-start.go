package Cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Daemon"
	"goTest/domain/Util"
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
		commandName := os.Args[2]
		scheduleList := *Daemon.RegisterDaemonSchedule()
		processItemMap, isExist := scheduleList[commandName]
		if !isExist {
			panic("找不到对应的进程名=" + commandName)
		}

		flagsFn, isExistFlagFn := processItemMap["flags"]
		if !isExistFlagFn {
			panic(commandName + "找不到对应的flags")
		}
		flags := flagsFn(cmd)
		log.Println(flags)
		parseFlag(cmd, flags)
	}
}

func startDaemon(cmd *cobra.Command, args []string) {
	commandName := args[0]
	scheduleList := *Daemon.RegisterDaemonSchedule()
	processItemMap, isExist := scheduleList[commandName]
	if !isExist {
		panic("找不到对应的进程名=" + commandName)
	}

	fn, isExistFn := processItemMap["fn"]
	if !isExistFn {
		panic(commandName + "找不到对应的处理函数")
	}

	if Util.ContainsInSlice(args, "d") || Util.ContainsInSlice(args, "D") {
		forkDaemonProcess(args)
		os.Exit(0)
	} else {
		startProcess(commandName, fn, cmd)
	}
}

func startProcess(commandName string, fn func(cmd *cobra.Command) []string, cmd *cobra.Command) {
	// 判断进程是否已经启动了
	//pid := getDaemonProcessPid(getPidFilePath(commandName))
	//if pid > 0 {
	//	if isProcessRunning(pid) {
	//		log.Printf("进程ID=%d已经启动，无需重新启动", pid)
	//		return
	//	}
	//}
	createDaemonPidPath()
	saveDaemonProcessPid(commandName)

	channel := make(chan int, 1)
	go func(channel chan int) {
		fn(cmd)
	}(channel)

	c := cron.New()
	_, _ = c.AddFunc("@every 2s", func() {
		saveDaemonProcessPid(commandName)
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
		newCmd.Wait()
		err := newCmd.Start()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "创建exec守护进程失败: %s\n", err)
		}
	}
}
