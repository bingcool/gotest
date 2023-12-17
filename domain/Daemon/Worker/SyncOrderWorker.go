package Worker

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func Run(cmd *cobra.Command) {

	for {
		fmt.Println("daemon run......")
		time.Sleep(1 * time.Second)
	}

}
