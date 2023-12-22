package Daemon

import (
	"github.com/spf13/cobra"
	"goTest/domain/Daemon/Worker"
)

type ScheduleFunc func(cmd *cobra.Command) []string
type ScheduleType map[string]map[string]ScheduleFunc

var daemonSchedule ScheduleType

func RegisterDaemonSchedule() *ScheduleType {
	daemonSchedule = ScheduleType{
		// 用户数据
		"consume-user-order": map[string]ScheduleFunc{
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
