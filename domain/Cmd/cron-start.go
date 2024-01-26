package Cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Cron"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// go run main.go cron-start --name=bingcool

var cronStartCommandName = "cron-start"

var CronStartCmd = &cobra.Command{
	Use:   cronStartCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		// log.Printf("daemon before run ")
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		if value, _ := cmd.Flags().GetString("fork_cron"); value == "no" {
			startCron(cmd)
		} else {
			fmt.Println("cron service daemon start")
			forkCronProcess(args)
			time.Sleep(1 * time.Second)
			os.Exit(0)
		}
	},
}

func init() {
	initCronStartFlags()
}

func initCronStartFlags() {
	if os.Args[1] == cronStartCommandName {
		if len(os.Args) > 2 {
			parseFlag(CronStartCmd, os.Args[2:])
		}
	}
}

func forkCronProcess(args []string) {
	osName := runtime.GOOS
	switch osName {
	// linux，macos
	case "linux", "darwin":
		newArgs := make([]string, 0)
		newArgs = append(newArgs, cronStartCommandName)
		for _, value := range args {
			if value != "d" && value != "D" {
				newArgs = append(newArgs, value)
			}
		}
		newArgs = append(newArgs, "--fork_cron=no")
		newCmd := exec.Command(os.Args[0], newArgs...)
		newCmd.Stdin = os.Stdin
		newCmd.Stdout = os.Stdout
		newCmd.Stderr = os.Stderr
		err := newCmd.Start()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "创建cron守护进程失败: %s\n", err)
		}
	}
}

func startCron(cmd *cobra.Command) {
	saveCronServerPid(os.Getpid())
	scheduleList := *Cron.RegisterCronSchedule()
	crontab := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
	for processName, scheduleItem := range scheduleList {
		specFn := scheduleItem["spec"]
		specSlice := specFn(cmd)
		for _, value := range specSlice {
			_, _ = crontab.AddFunc(value, func() {
				execDaemonProcess(processName)
			})
		}
		time.Sleep(100 * time.Microsecond)
	}
	_, _ = crontab.AddFunc("@every 10s", func() {
		saveCronServerPid(os.Getpid())
	})
	crontab.Start()
	fmt.Println("cron service start.")
	select {}
}

func execDaemonProcess(processName string) {
	osName := runtime.GOOS
	switch osName {
	// linux，macos
	case "linux", "darwin":
		newArgs := make([]string, 0)
		//调用拉起新守护进程处理定时任务
		newArgs = append(newArgs, daemonStartCommandName, processName)
		newArgs = append(newArgs, "--from-flag=cron")
		newCmd := exec.Command(os.Args[0], newArgs...)
		newCmd.Stdin = os.Stdin
		newCmd.Stdout = os.Stdout
		newCmd.Stderr = os.Stderr
		err := newCmd.Start()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "创建cron守护进程失败: %s\n", err)
		}
	}
}

func cronHandle(processName string, fn func(cmd *cobra.Command) []string, cmd *cobra.Command) {
	createCronPidPath()
	saveProcessPid(getCronPidFile(processName))
	fn(cmd)
}
