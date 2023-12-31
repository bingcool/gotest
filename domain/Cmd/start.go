package Cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Middlewares"
	"goTest/domain/Router"
	"goTest/domain/System"
	"goTest/domain/Util"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// go run main.go start --myname=bingcool

var startCommandName = "start"

var StartCmd = &cobra.Command{
	Use:   startCommandName,
	Short: "start the gofy",
	Long:  `start the gofy`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		fmt.Println("before run ")
	},

	Run: func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之后执行的操作
		fmt.Println("after run")
	},
}

func init() {
	initStartParseFlag()
}

func initStartParseFlag() {
	if os.Args[1] == startCommandName {
		if len(os.Args) > 2 {
			parseFlag(StartCmd, os.Args[2:])
		}
	}
}

func run(cmd *cobra.Command, args []string) {
	Console.NewConsole().PutCommand(cmd)
	if Util.ContainsInSlice(args, "d") || Util.ContainsInSlice(args, "D") {
		osName := runtime.GOOS
		switch osName {
		// linux，macos
		case "linux", "darwin":
			newArgs := make([]string, 0)
			newArgs = append(newArgs, startCommandName)
			for _, value := range args {
				if value != "d" && value != "D" {
					newArgs = append(newArgs, value)
				}
			}
			newCmd := exec.Command(os.Args[0], newArgs...)
			newCmd.Stdin = os.Stdin
			//cmd1.Stderr = os.Stderr
			err := newCmd.Start()
			if err != nil {
				_, err := fmt.Fprintf(os.Stderr, "[-] Error: %s\n", err)
				if err != nil {
					os.Exit(0)
				}
			}
			os.Exit(0)
		default:
			startServer()
		}
	} else {
		startServer()
	}
}

// StartServer 解析
func startServer() {
	r := gin.Default()
	// 设置日志中间件，主要用于打印请求日志
	// 设置Recovery中间件，主要用于拦截panic错误，不至于导致进程崩掉
	Middlewares.SetGlobalMiddleware(r)
	Router.SetupProductRouter(r)
	Router.SetOrderRouter(r)
	System.SaveMainPid()
	crontab := cron.New()
	_, _ = crontab.AddFunc("@every 10s", func() {
		System.SaveMainPid()
	})

	err := r.Run(":8080")
	// 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		log.Fatal("启动失败")
	}
}
