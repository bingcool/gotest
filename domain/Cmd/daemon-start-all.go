package Cmd

import (
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Daemon"
	"log"
	"os"
	"time"
)

var daemonStartAllCommandName = "daemon-start-all"

var DaemonStartAllCmd = &cobra.Command{
	Use:   daemonStartAllCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		// log.Printf("daemon before run ")
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		startDamonAll()
	},
}

func init() {
	initParseFlag(daemonStartAllCommandName, DaemonStartAllCmd)
}

func startDamonAll() {
	scheduleList := *Daemon.RegisterDaemonSchedule()
	for commandName := range scheduleList {
		log.Println(commandName)
		newArgs := make([]string, 0)
		newArgs = append(newArgs, commandName)
		forkDaemonProcess(newArgs)
		log.Printf("启动进程【%s】", commandName)
		time.Sleep(100 * time.Microsecond)
	}
	os.Exit(0)
}
