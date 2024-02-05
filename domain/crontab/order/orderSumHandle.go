package order

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func Run(cmd *cobra.Command) {
	name, _ := cmd.Flags().GetString("name")
	fmt.Println("crontab run......", name)
	time.Sleep(1 * time.Second)
}

func Run1(cmd *cobra.Command) {

	for {
		fmt.Println("crontab run1111......")
		time.Sleep(1 * time.Second)
	}

}
