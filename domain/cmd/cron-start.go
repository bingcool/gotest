package cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"goTest/domain/crontab"
	"goTest/domain/system"
	"os"
	"os/exec"
	"time"
)

// go run main.go crontab-start --name=bingcool

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
		console.NewConsole().PutCommand(cmd)
		if value, _ := cmd.Flags().GetInt("daemon"); value == 0 {
			startCron(cmd)
		} else {
			//fmt.Println("crontab service daemon start")
			//time.Sleep(2 * time.Second)

			fmt.Println("crontab service daemon start")
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
	initDaemonFlags(CronStartCmd)
}

func forkCronProcess(args []string) {
	if system.IsLinux() || system.IsMacos() {
		newArgs := make([]string, 0)
		newArgs = append(newArgs, cronStartCommandName)
		for _, value := range args {
			if value != "d" && value != "D" {
				newArgs = append(newArgs, value)
			}
		}
		newArgs = append(newArgs, "--daemon=0")
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

// 定时任务调度池
var crontabSchedulePool []*cron.Cron

func startCron(cmd *cobra.Command) {
	saveCronServerPid(os.Getpid())
	scheduleList := *crontab.RegisterCronSchedule()
	for processName, scheduleItem := range scheduleList {
		specFn := scheduleItem["spec"]
		specSlice := specFn(cmd)
		for _, value := range specSlice {
			crontabSchedule := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
			_, _ = crontabSchedule.AddFunc(value, func() {
				fmt.Println(processName)
				execDaemonProcess(processName)
			})
			crontabSchedule.Start()
			crontabSchedulePool = append(crontabSchedulePool, crontabSchedule)
		}
		time.Sleep(100 * time.Microsecond)
	}

	mainCrontabSchedule := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
	_, _ = mainCrontabSchedule.AddFunc("@every 2s", func() {
		saveCronServerPid(os.Getpid())
	})
	mainCrontabSchedule.Start()

	crontabSchedulePool = append(crontabSchedulePool, mainCrontabSchedule)
	fmt.Println("crontab service start.")
	select {}
}

// execDaemonProcess 任务时间到，拉起新的进程处理任务
func execDaemonProcess(processName string) {
	if system.IsLinux() || system.IsMacos() {
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
			_, _ = fmt.Fprintf(os.Stderr, "execDaemonProcess创建cron守护进程失败: %s\n", err)
		}
	}
}

func cronHandle(processName string, fn func(cmd *cobra.Command) []string, cmd *cobra.Command) {
	createCronPidPath()
	saveProcessPid(getCronPidFile(processName))
	fn(cmd)
}
