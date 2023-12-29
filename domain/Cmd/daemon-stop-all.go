package Cmd

import (
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"log"
	"os"
	"path/filepath"
)

var daemonStopAllCommandName = "daemon-stop-all"

var DaemonStopAllCmd = &cobra.Command{
	Use:   daemonStopAllCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		// log.Printf("daemon before run ")
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		stopDaemonAll(true)
	},
}

func init() {
	initParseFlag(daemonStopAllCommandName, DaemonStopAllCmd)
}

func stopDaemonAll(isExit bool) {
	pidFilePath := getDaemonPidPath()
	files, err := os.ReadDir(pidFilePath)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	for _, file := range files {
		pidFile := filepath.Join(pidFilePath, file.Name())
		pid := getProcessPid(pidFile)
		killProcess(pid, pidFile)
	}

	if isExit {
		os.Exit(0)
	}
}
