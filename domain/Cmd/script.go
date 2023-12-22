package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Script"
)

var scriptCommand = "script"

var ScriptCmd = &cobra.Command{
	Use:   scriptCommand,
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
	initParseFlag(scriptCommand, ScriptCmd)
}
