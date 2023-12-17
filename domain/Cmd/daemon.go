package Cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"goTest/domain/Console"
	"goTest/domain/Daemon"
	"goTest/domain/System"
	"goTest/domain/Util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

var daemonCommand = "daemon"

var DaemonCmd = &cobra.Command{
	Use:   daemonCommand,
	Short: "run script",
	Long:  "run script",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 在每个命令执行之前执行的操作
		log.Printf("daemon before run ")
	},
	Run: func(cmd *cobra.Command, args []string) {
		Console.NewConsole().PutCommand(cmd)
		action := args[0]
		switch action {
		case "restartAll":
			restartAll(cmd, args)
		case "start":
			startDaemon(cmd, args)
		case "stop":
			stopDaemon(cmd, args)
		}
	},
}

func init() {
	ParseFlags(daemonCommand, DaemonCmd)
}

func restartAll(cmd *cobra.Command, args []string) {
	pidFilePath := getDaemonPidPath()
	files, err := os.ReadDir(pidFilePath)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	for _, file := range files {
		pidFile := filepath.Join(pidFilePath, file.Name())
		pid := getDaemonProcessPid(pidFile)
		killProcess(pid, pidFile)
	}

	time.Sleep(2 * time.Second)

	schedule := *Daemon.GetDaemonSchedule()
	for commandName, _ := range schedule {
		log.Println(commandName)
		//args[1] = commandName
		//forkDaemonProcess(args)
		//log.Printf("启动进程【%s】", commandName)
		//time.Sleep(100 * time.Microsecond)
	}
}

func startDaemon(cmd *cobra.Command, args []string) {
	commandName := args[1]
	schedule := *Daemon.GetDaemonSchedule()
	fn, isExist := schedule[commandName]
	if !isExist {
		panic("找不到对应的进程名=" + commandName)
	}

	if Util.ContainsInSlice(args, "d") || Util.ContainsInSlice(args, "D") {
		forkDaemonProcess(args)
	} else {
		startProcess(commandName, fn, cmd)
	}
}

func startProcess(commandName string, fn func(cmd *cobra.Command), cmd *cobra.Command) {
	// 判断进程是否已经启动了
	pid := getDaemonProcessPid(getPidFilePath(commandName))
	if pid > 0 {
		if isProcessRunning(pid) {
			log.Printf("进程ID=%d已经启动，无需重新启动", pid)
			return
		}
	}

	saveDaemonProcessPid(commandName)
	channel := make(chan int, 1)
	go func(channel chan int) {
		fn(cmd)
	}(channel)

	c := cron.New()
	_, _ = c.AddFunc("@every 2s", func() {
		saveDaemonProcessPid(commandName)
	})
	c.Start()

	select {
	case <-channel:
	}
}

func stopDaemon(cmd *cobra.Command, args []string) {
	commandName := args[1]
	pidFile := filepath.Join(getDaemonPidPath(), commandName+".pid")
	pid := getDaemonProcessPid(pidFile)
	killProcess(pid, pidFile)
}

func killProcess(pid int, filePath string) {
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Printf("Error finding process:", err)
		if len(filePath) > 0 {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			log.Printf("File deleted successfully")
		}
	} else {
		_ = process.Signal(syscall.SIGTERM)
		log.Printf("process stop successfully")
	}
}

func isProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		log.Printf("Process %d is dead!", pid)
		return false
	} else {
		log.Printf("Process %d is alive!", pid)
		return true
	}
}

func saveDaemonProcessPid(commandName string) {
	pidFile := getPidFilePath(commandName)
	pid := os.Getpid()
	serverFile, _ := os.Create(pidFile)
	_, err := serverFile.WriteString(strconv.Itoa(pid))
	if err != nil {
		log.Fatal("save server.pid error")
	}
}

func getPidFilePath(commandName string) string {
	pidFile := filepath.Join(getDaemonPidPath(), commandName+".pid")
	return pidFile
}

func getDaemonProcessPid(pidFilePath string) int {
	pid, err := os.ReadFile(pidFilePath)
	if err != nil {
		return 0
	}
	var serverPid int
	serverPid, _ = strconv.Atoi(string(pid))
	return serverPid
}

func getDaemonPidPath() string {
	return filepath.Join(System.GetDaemonPath(), "Pid")
}

func forkDaemonProcess(args []string) {
	osName := runtime.GOOS
	switch osName {
	// linux，macos
	case "linux", "darwin":
		newArgs := make([]string, 0)
		//
		newArgs = append(newArgs, daemonCommand)
		for _, value := range args {
			if value != "d" && value != "D" {
				newArgs = append(newArgs, value)
			}
		}
		newCmd := exec.Command(os.Args[0], newArgs...)
		newCmd.Stdin = os.Stdin
		//newCmd.Stdout = os.Stdout
		newCmd.Stderr = os.Stderr
		err := newCmd.Start()
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "[-] Error: %s\n", err)
			if err != nil {
				os.Exit(0)
			}
		}
		os.Exit(0)
	}
}
