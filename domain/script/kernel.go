package script

import (
	"github.com/spf13/cobra"
	"goTest/domain/script/command"
)

var scriptSchedule map[string]func(cmd *cobra.Command)

func RegisterScriptSchedule() *map[string]func(cmd *cobra.Command) {
	scriptSchedule = map[string]func(cmd *cobra.Command){
		// 修复用户数据
		"fix-user": func(cmd *cobra.Command) {
			command.Test1(cmd)
		},

		// 修复用户数据
		"fix-user1": func(cmd *cobra.Command) {
			command.Test1(cmd)
		},
	}
	return &scriptSchedule
}
