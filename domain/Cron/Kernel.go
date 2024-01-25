package Cron

import (
	"goTest/domain/Console"
	"goTest/domain/Cron/Conf"
)

var cronSchedule Console.ScheduleType

func RegisterCronSchedule() *Console.ScheduleType {
	cronSchedule = Console.ScheduleType{
		// 用户数据
		"cron-user-order": Conf.RegisterUserOrder(),
	}

	return &cronSchedule
}
