package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Util"
	"os"
	"strconv"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "gofy",
	Short: "My Gin application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init running")
		run(cmd, args)
	},
}

// Env 全局flags
var (
	Env string
)

func init() {
	rootCmd.AddCommand(StartCmd)
	rootCmd.AddCommand(StopCmd)
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(ScriptCmd)
	rootCmd.AddCommand(DaemonStartCmd)
	rootCmd.AddCommand(DaemonStartAllCmd)
	rootCmd.AddCommand(DaemonStopCmd)
	rootCmd.AddCommand(DaemonStopAllCmd)
	rootCmd.AddCommand(CronStartCmd)

	StartCmd.Flags().StringVar(&Env, "environment", "dev", "environment of system")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// ParseFlags
func initParseFlag(useCmd string, runCmd *cobra.Command) {
	args := os.Args
	if len(args) >= 3 && args[1] == useCmd {
		flags := args[3:]
		parseFlag(runCmd, flags)
	}
}

func parseFlag(runCmd *cobra.Command, flags []string) {
	charsToRemove := "-"
	for _, v := range flags {
		item := strings.Split(v, "=")
		if len(item) != 2 {
			continue
		}
		var flagNameString string
		var flagNameInt int
		var flagNameFloat float64
		flagName := strings.ReplaceAll(item[0], charsToRemove, "")
		flagValue := item[1]
		if Util.IsNumber(flagValue) {
			if Util.IsInt(flagValue) {
				flagValueNumber, _ := strconv.Atoi(flagValue)
				runCmd.Flags().IntVar(&flagNameInt, flagName, flagValueNumber, "int flags params")
			} else if Util.IsFloat(flagValue) {
				flagValueFloat, _ := strconv.ParseFloat(flagValue, 64)
				runCmd.Flags().Float64Var(&flagNameFloat, flagName, flagValueFloat, "float flags params")
			}
		} else {
			runCmd.Flags().StringVar(&flagNameString, flagName, flagValue, "string flags params")
		}
	}
}
