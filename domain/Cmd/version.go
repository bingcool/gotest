package Cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
)

var version = "1.0.0"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show version",
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		fmt.Println(version)
	},
}
