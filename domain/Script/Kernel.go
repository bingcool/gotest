package Script

import (
	"github.com/spf13/cobra"
	"goTest/domain/Script/Command"
)

var schedule map[string]func(cmd *cobra.Command)

func GetSchedule() *map[string]func(cmd *cobra.Command) {
	schedule = map[string]func(cmd *cobra.Command){
		"fix-user": func(cmd *cobra.Command) {
			Command.Test1(cmd)
		},

		"fix-user1": func(cmd *cobra.Command) {
			Command.Test1(cmd)
		},
	}
	return &schedule
}
