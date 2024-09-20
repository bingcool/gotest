package daemon

import (
	"goTest/domain/console"
	"goTest/domain/daemon/conf"
)

var daemonSchedule console.ScheduleType

func RegisterDaemonSchedule() *console.ScheduleType {
	daemonSchedule = console.ScheduleType{
		// 用户数据
		"consume-user-orderDto": conf.RegisterSyncUserOrder(),
	}
	return &daemonSchedule
}
