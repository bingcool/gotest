package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/util"
	"strconv"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "gofy",
	Short: "My Gin application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root running")
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
	rootCmd.AddCommand(CronStopCmd)

	StartCmd.Flags().StringVar(&Env, "env", "dev", "environment of system")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func parseFlag(runCmd *cobra.Command, flags []string) {
	charsToRemove := "-"
	for _, v := range flags {
		item := strings.Split(v, "=")
		if len(item) != 2 {
			continue
		}
		flagName := strings.Replace(item[0], charsToRemove, "", 2)
		flagValue := item[1]
		// flag已存在则跳过
		if runCmd.Flags().Lookup(flagName) != nil {
			continue
		}

		var flagNameString string
		var flagNameInt int
		var flagNameFloat float64

		if util.IsNumber(flagValue) {
			if util.IsInt(flagValue) {
				flagValueNumber, _ := strconv.Atoi(flagValue)
				runCmd.Flags().IntVar(&flagNameInt, flagName, flagValueNumber, "int flags params")
			} else if util.IsFloat(flagValue) {
				flagValueFloat, _ := strconv.ParseFloat(flagValue, 64)
				runCmd.Flags().Float64Var(&flagNameFloat, flagName, flagValueFloat, "float flags params")
			}
		} else {
			runCmd.Flags().StringVar(&flagNameString, flagName, flagValue, "string flags params")
		}
	}
}
