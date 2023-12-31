package Daemon

import (
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Daemon/Worker"
)

var daemonSchedule Console.ScheduleType

func RegisterDaemonSchedule() *Console.ScheduleType {
	daemonSchedule = Console.ScheduleType{
		// 用户数据
		"consume-user-order": map[string]Console.ScheduleFunc{
			"flags": func(cmd *cobra.Command) []string {
				flags := make([]string, 0)
				flags = append(flags, "--name=huang")
				return flags
			},
			"fn": func(cmd *cobra.Command) []string {
				Worker.Run(cmd)
				return make([]string, 0)
			},
		},

		//// 用户数据
		//"consume-user-order1": map[string]any{
		//	"params": "kkkkk",
		//	"fn": func(cmd *cobra.Command) {
		//		Worker.Run1(cmd)
		//	},
		//},
	}
	return &daemonSchedule
}
