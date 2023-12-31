package Cron

import (
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Cron/Order"
)

var cronSchedule Console.ScheduleType

func RegisterCronSchedule() *Console.ScheduleType {
	cronSchedule = Console.ScheduleType{
		// 用户数据
		"cron-user-order": map[string]Console.ScheduleFunc{
			"spec": func(cmd *cobra.Command) []string {
				spec := make([]string, 0)
				spec = append(spec, "@every 10s")
				return spec
			},
			"flags": func(cmd *cobra.Command) []string {
				flags := make([]string, 0)
				flags = append(flags, "--name=huang")
				return flags
			},
			"fn": func(cmd *cobra.Command) []string {
				Order.Run(cmd)
				return make([]string, 0)
			},
		},

		//// 用户数据
		//"consume-user-order1": map[string]any{
		//	"params": "kkkkk",
		//	"fn": func(cmd *cobra.Command) {
		//		Order.Run1(cmd)
		//	},
		//},
	}
	return &cronSchedule
}
