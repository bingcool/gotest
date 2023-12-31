package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Script"
	"os"
)

// go run main.go script fix-user --order_id=11111

var scriptCommandName = "script"

var ScriptCmd = &cobra.Command{
	Use:   scriptCommandName,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		fmt.Println("script before run ")
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		schedule := *Script.RegisterScriptSchedule()
		callFunc := schedule[args[0]]
		callFunc(cmd)
	},
}

func init() {
	initScriptParseFlag()
}

func initScriptParseFlag() {
	if os.Args[1] == scriptCommandName {
		if len(os.Args) > 3 {
			parseFlag(ScriptCmd, os.Args[3:])
		}
	}
}
