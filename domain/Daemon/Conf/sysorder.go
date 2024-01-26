package Conf

import (
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Daemon/Worker"
)

func RegisterSyncUserOrder() Console.ScheduleMap {
	schedule := Console.NewSchedule()

	schedule.AddFlags(func(cmd *cobra.Command) []string {
		flags := make([]string, 0)
		flags = append(flags, "--name=huang")
		return flags
	})

	schedule.AddFn(func(cmd *cobra.Command) []string {
		Worker.Run(cmd)
		return make([]string, 0)
	})

	return schedule.ScheduleMap
}
