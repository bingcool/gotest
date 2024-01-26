package Daemon

import (
	"goTest/domain/Console"
	"goTest/domain/Daemon/Conf"
)

var daemonSchedule Console.ScheduleType

func RegisterDaemonSchedule() *Console.ScheduleType {
	daemonSchedule = Console.ScheduleType{
		// 用户数据
		"consume-user-order": Conf.RegisterSyncUserOrder(),
	}
	return &daemonSchedule
}
