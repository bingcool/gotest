package conf

import (
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"goTest/domain/crontab/order"
)

func RegisterUserOrder() console.ScheduleMap {
	schedule := console.NewSchedule()

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
		order.Run(cmd)
		return make([]string, 0)
	})

	return schedule.ScheduleMap
}
