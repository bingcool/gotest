package crontab

import (
	"goTest/domain/console"
	"goTest/domain/crontab/conf"
)

var cronSchedule console.ScheduleType

func RegisterCronSchedule() *console.ScheduleType {
	cronSchedule = console.ScheduleType{
		// 用户数据
		"crontab-user-order": conf.RegisterUserOrder(),
	}

	return &cronSchedule
}
