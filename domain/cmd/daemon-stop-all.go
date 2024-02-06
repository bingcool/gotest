package cmd

import (
	"flag"
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"log"
	"os"
	"path/filepath"
)

// go run main.go daemon-stop-all

var daemonStopAllCommandName = "daemon-stop-all"

var DaemonStopAllCmd = &cobra.Command{
	Use:   daemonStopAllCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		console.NewConsole().PutCommand(cmd)
		stopDaemonAll(true)
	},
}

func init() {
	initDaemonStopAllFlag()
}

func initDaemonStopAllFlag() {
	if os.Args[1] == daemonStopAllCommandName {
		if len(os.Args) > 2 {
			parseFlag(DaemonStopAllCmd, flag.Args()[2:])
		}
	}
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
