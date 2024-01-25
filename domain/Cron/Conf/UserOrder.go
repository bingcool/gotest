package Conf

import (
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Cron/Order"
)

func RegisterUserOrder() Console.ScheduleMap {
	schedule := Console.NewSchedule()

	schedule.AddSpec(func(cmd *cobra.Command) []string {
		spec := make([]string, 0)
		spec = append(spec, "@every 10s")
		return spec
	})

	schedule.AddFlags(func(cmd *cobra.Command) []string {
		flags := make([]string, 0)
		flags = append(flags, "--name=huang")
		return flags
	})

	schedule.AddFn(func(cmd *cobra.Command) []string {
		Order.Run(cmd)
		return make([]string, 0)
	})

	return schedule.ScheduleMap
}
