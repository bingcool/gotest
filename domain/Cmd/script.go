package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Script"
	"goTest/domain/Util"
	"os"
	"strconv"
	"strings"
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
		schedule := Script.GetSchedule()
		callFunc := (*schedule)[args[0]]
		callFunc(cmd)
	},
}

func init() {
	parseFlags()
}

func parseFlags() {
	args := os.Args
	charsToRemove := "-"
	if len(args) >= 2 && args[1] == scriptCommand {
		for _, v := range args[3:] {
			item := strings.Split(v, "=")
			var flagNameString string
			var flagNameInt int
			var flagNameFloat float64
			flagName := strings.ReplaceAll(item[0], charsToRemove, "")
			flagValue := item[1]
			if Util.IsNumber(flagValue) {
				if Util.IsInt(flagValue) {
					flagValueNumber, _ := strconv.Atoi(flagValue)
					ScriptCmd.Flags().IntVar(&flagNameInt, flagName, flagValueNumber, "int flags params")
				} else if Util.IsFloat(flagValue) {
					flagValueFloat, _ := strconv.ParseFloat(flagValue, 64)
					ScriptCmd.Flags().Float64Var(&flagNameFloat, flagName, flagValueFloat, "float flags params")
				}
			} else {
				ScriptCmd.Flags().StringVar(&flagNameString, flagName, flagValue, "string flags params")
			}
		}
	}

}
