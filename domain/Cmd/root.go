package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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

	StartCmd.Flags().StringVar(&Env, "environment", "dev", "environment of system")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
