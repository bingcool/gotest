package Order

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func Run(cmd *cobra.Command) {
	name, _ := cmd.Flags().GetString("name")
	fmt.Println("cron run......", name)
	time.Sleep(1 * time.Second)
}

func Run1(cmd *cobra.Command) {

	for {
		fmt.Println("cron run1111......")
		time.Sleep(1 * time.Second)
	}

}
