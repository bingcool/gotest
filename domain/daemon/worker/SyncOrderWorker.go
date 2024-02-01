package worker

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func Run(cmd *cobra.Command) {

	for {
		name, _ := cmd.Flags().GetString("name")
		fmt.Println("daemon run......", name)
		time.Sleep(1 * time.Second)
	}

}

func Run1(cmd *cobra.Command) {

	for {
		fmt.Println("daemon run1111......")
		time.Sleep(1 * time.Second)
	}

}
