package Cmd

import (
	"flag"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Daemon"
	"log"
	"os"
	"time"
)

// go run main.go daemon-start-all

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
	initDaemonStartAllFlags()
}

func initDaemonStartAllFlags() {
	if os.Args[1] == daemonStartAllCommandName {
		if len(os.Args) > 2 {
			parseFlag(DaemonStartAllCmd, flag.Args()[2:])
		}
	}
}

func startDamonAll() {
	scheduleList := *Daemon.RegisterDaemonSchedule()
	for processName := range scheduleList {
		log.Println(processName)
		newArgs := make([]string, 0)
		newArgs = append(newArgs, processName)
		forkDaemonProcess(newArgs)
		log.Printf("启动进程【%s】", processName)
		time.Sleep(100 * time.Microsecond)
	}
	os.Exit(0)
}
