package Cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Cron"
	"log"
	"os"
	"time"
)

var cronCommandName = "cron"

var CronStartCmd = &cobra.Command{
	Use:   cronCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		// log.Printf("daemon before run ")
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		startCron(cmd, args)
	},
}

func init() {
	initParseFlag(cronCommandName, CronStartCmd)
}

func startCron(cmd *cobra.Command, args []string) {
	scheduleList := *Cron.RegisterCronSchedule()
	for commandName := range scheduleList {
		log.Println(commandName)
		newArgs := make([]string, 0)
		newArgs = append(newArgs, commandName)
		forkDaemonProcess(newArgs)
		log.Printf("启动进程【%s】", commandName)
		time.Sleep(100 * time.Microsecond)
	}

	os.Exit(0)
	c := cron.New()
	fmt.Println("gogoogoogoo")
	// 添加定时任务
	c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("每5分钟执行一次")
	})
	//c.AddFunc("@hourly", func() {
	//	fmt.Println("每小时执行一次")
	//})

	// 启动定时任务
	c.Start()

	// 程序运行一段时间后停止定时任务
	// time.Sleep(1 * time.Hour)
	// c.Stop()

	// 阻塞主协程，保持程序运行
	select {}
}
