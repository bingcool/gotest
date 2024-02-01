package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"os"
	"syscall"
)

// go run main.go crontab-stop

var cronStopCommandName = "cron-stop"

func init() {
	initCronStopParseFlag()
}

func initCronStopParseFlag() {
	if os.Args[1] == cronStopCommandName {
		if len(os.Args) > 2 {
			parseFlag(StartCmd, os.Args[2:])
		}
	}
}

var CronStopCmd = &cobra.Command{
	Use:   cronStopCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		console.NewConsole().PutCommand(cmd)
		cronServerPid := getCronServerPid()
		if isProcessRunning(cronServerPid) {
			process, err := os.FindProcess(cronServerPid)
			if err != nil {
				fmt.Println("Error finding process:", err)
				return
			}
			fmt.Println("crontab service stop")
			_ = process.Signal(syscall.SIGTERM)
		}
	},
}
