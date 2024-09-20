package main

import (
	"goTest/domain/cmd"
	"goTest/domain/system"
)

func init() {
	//fmt.Println("Main package init")
}

func main() {

	system.GetRootDir()
	_ = cmd.Execute()
}
