package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "goTest/docs"
	"goTest/domain/console"
	"goTest/domain/middlewares"
	"goTest/domain/route"
	"goTest/domain/system"
	"log"
	"os"
	"os/exec"
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
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {

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

		initDaemonFlags(StartCmd)
	}
}

func run(cmd *cobra.Command, args []string) {
	console.NewConsole().PutCommand(cmd)
	if isFork(cmd) && (system.IsLinux() || system.IsMacos()) {
		newArgs := make([]string, 0)
		newArgs = append(newArgs, startCommandName)
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
	} else {
		startServer()
	}
}

// StartServer 解析
func startServer() {
	r := gin.Default()
	// 设置Recovery中间件，主要用于拦截panic错误，不至于导致进程崩掉
	middlewares.SetGlobalMiddleware(r)
	route.SetupProductRouter(r)
	route.SetOrderRouter(r)
	system.SaveMainPid()
	crontab := cron.New()
	_, _ = crontab.AddFunc("@every 10s", func() {
		system.SaveMainPid()
	})

	// swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := r.Run(":8080")
	// 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		log.Fatal("启动失败")
	}
}
