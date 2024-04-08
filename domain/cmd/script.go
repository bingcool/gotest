package cmd

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"
	"goTest/domain/console"
	"goTest/domain/script"
	"log"
	"os"
	"time"
)

// go run main.go script fix-user --order_id=11111

var scriptCommandName = "script"

var ScriptCmd = &cobra.Command{
	Use:   scriptCommandName,
	Short: "run script",
	Long:  "run script",
	Args:  cobra.MaximumNArgs(1), //使用内置的验证函数，位置参数只能一个，即命令之后的变量，这里是指脚本名称
	// 如果设置了PersistentPreRun，将会覆盖rootCmd设置的PersistentPreRun
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		fmt.Println("script PersistentPreRun")
		if len(os.Args) <= 2 {
			log.Fatal("请指定脚本执行的命令名称")
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		count := 10

		// create and start new bar
		bar := pb.StartNew(count)

		// start bar from 'default' template
		// bar := pb.Default.Start(count)

		// start bar from 'simple' template
		// bar := pb.Simple.Start(count)

		// start bar from 'full' template
		// bar := pb.Full.Start(count)

		for i := 0; i < count; i++ {
			bar.Increment()
			time.Sleep(time.Second)
		}

		// finish bar
		bar.Finish()
	},
	Run: func(cmd *cobra.Command, args []string) {
		console.NewConsole().PutCommand(cmd)
		scheduleList := *script.RegisterScriptSchedule()
		callFunc := scheduleList[args[0]]
		callFunc(cmd)
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
	// 如果设置了PersistentPostRun，将会覆盖rootCmd设置的PersistentPostRun
	PersistentPostRun: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	initScriptParseFlag()
}

func initScriptParseFlag() {
	if os.Args[1] == scriptCommandName {
		if len(os.Args) > 3 {
			parseFlag(ScriptCmd, os.Args[3:])
		}
	}

	//var Region string
	//// 默认情况下，标志是可选的。我们可以将其标记为必选，如果没有提供，则会报错。
	//ScriptCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	//// 设置为必须
	//_ = ScriptCmd.MarkFlagRequired("region")

}
