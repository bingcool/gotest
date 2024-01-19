package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"os"
	"syscall"
)

// go run main.go cron-stop

var cronStopCommandName = "cron-stop"

var CronStopCmd = &cobra.Command{
	Use:   cronStopCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		cronServerPid := getCronServerPid()
		if isProcessRunning(cronServerPid) {
			process, err := os.FindProcess(cronServerPid)
			if err != nil {
				fmt.Println("Error finding process:", err)
				return
			}
			fmt.Println("cron service stop")
			_ = process.Signal(syscall.SIGTERM)
		}
	},
}
