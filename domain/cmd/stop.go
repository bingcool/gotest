package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"goTest/domain/system"
	"os"
	"syscall"
)

// go run main.go stop

var stopCommandName = "stop"

var StopCmd = &cobra.Command{
	Use:   stopCommandName,
	Short: "stop the gofy",
	Long:  "stop the gofy",
	Run: func(cmd *cobra.Command, args []string) {
		console.NewConsole().PutCommand(cmd)
		// 将 PID 转换为 os.Process 类型
		process, err := os.FindProcess(system.GetMainPid())
		if err != nil {
			fmt.Println("Error finding process:", err)
			return
		}
		_ = process.Signal(syscall.SIGTERM)
	},
}

func init() {
	initStopParseFlag()
}

func initStopParseFlag() {
	if os.Args[1] == stopCommandName {
		if len(os.Args) > 2 {
			parseFlag(StopCmd, os.Args[2:])
		}
	}
}
