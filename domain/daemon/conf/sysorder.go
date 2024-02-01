package conf

import (
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"goTest/domain/daemon/worker"
)

func RegisterSyncUserOrder() console.ScheduleMap {
	schedule := console.NewSchedule()

	schedule.AddFlags(func(cmd *cobra.Command) []string {
		flags := make([]string, 0)
		//flags = append(flags, "--name=huang")
		return flags
	})

	schedule.AddFn(func(cmd *cobra.Command) []string {
		worker.Run(cmd)
		return make([]string, 0)
	})

	return schedule.ScheduleMap
}
