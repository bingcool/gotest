package Daemon

import (
	"github.com/spf13/cobra"
	"goTest/domain/Daemon/Worker"
)

var daemonSchedule map[string]func(cmd *cobra.Command)

func GetDaemonSchedule() *map[string]func(cmd *cobra.Command) {
	daemonSchedule = map[string]func(cmd *cobra.Command){
		// 用户数据
		"consume-user-order": func(cmd *cobra.Command) {
			Worker.Run(cmd)
		},
	}
	return &daemonSchedule
}
