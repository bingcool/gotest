package Command

import (
	"fmt"
	"github.com/spf13/cobra"
	"goTest/domain/library/order"
)

func Test1(cmd *cobra.Command) {
	fmt.Println("script test1......")

	name, _ := cmd.Flags().GetString("myname")
	age, _ := cmd.Flags().GetInt("myage")
	kk, _ := cmd.Flags().GetString("kk")
	price, _ := cmd.Flags().GetFloat64("price")

	fmt.Println(name, age, kk, price)

	orderService := order.NewOrderService()
	list := orderService.GetOrderList(1691463646)
	fmt.Println(list)
}
