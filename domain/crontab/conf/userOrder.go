package conf

import (
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"goTest/domain/crontab/order"
)

func RegisterUserOrder() console.ScheduleMap {
	schedule := console.NewSchedule()

	// 定时任务时间格式
	schedule.AddSpec(func(cmd *cobra.Command) []string {
		spec := make([]string, 0)
		spec = append(spec, "@every 10s")
		return spec
	})

	// 定时任务启动内置的flag,可以通过flag值做不同逻辑，比如同步大量数据数据，flag=1时，同步1-100000，flag=2时，同步100000+
	schedule.AddFlags(func(cmd *cobra.Command) []string {
		flags := make([]string, 0)
		flags = append(flags, "--name=huang")
		return flags
	})

	// 定时任务处理函数
	schedule.AddFn(func(cmd *cobra.Command) []string {
		order.Run(cmd)
		return make([]string, 0)
	})

	return schedule.ScheduleMap
}
