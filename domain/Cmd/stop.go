package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/System"
	"os"
	"syscall"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop the gofy",
	Long:  "stop the gofy",
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		// 将 PID 转换为 os.Process 类型
		process, err := os.FindProcess(System.GetMainPid())
		if err != nil {
			fmt.Println("Error finding process:", err)
			return
		}
		_ = process.Signal(syscall.SIGTERM)
	},
}

func init() {

}
