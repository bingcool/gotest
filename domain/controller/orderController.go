package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goTest/domain/Go"
	"goTest/domain/dto/orderDto"
	LibraryOrder "goTest/domain/library/order"
)

// OrderController 定义结构体
type OrderController struct {
	BaseController
	OrderStatus int
}

// NewOrder 实体对象
func NewOrder(c *gin.Context) *OrderController {
	// 初始化函数
	order := &OrderController{}
	order.BaseController.Construct(c)
	return order
}

// ListOrder 订单列表

func (Order *OrderController) ListOrder(ctx *gin.Context, listOrderReqDto *orderDto.ListOrderReqDto) {
	fmt.Println(listOrderReqDto.OrderId)

	orderService := LibraryOrder.NewOrderService()
	result := orderService.GetOrderList(1691463954)
	var list []orderDto.ListOrderDto
	// list := make([]ListOrderDto, 0)
	i := 0
	for _, item := range result {
		listOrderDto := orderDto.ListOrderDto{}
		listOrderDto.OrderId = item.OrderId
		listOrderDto.UserId = item.UserId
		listOrderDto.JsonData = item.JsonData
		address := map[string]any{
			"country": "中国",
			"city":    "深圳",
		}
		listOrderDto.Address = address
		list = append(list, listOrderDto)
		i++
	}

	//m := map[string]any{"name": "bingcool"}
	//ch := make(chan any)
	Go.Run(func(ctx Go.Context) {
		//fmt.Println(ctx.BucketMap["name"])
	}, Go.Context{})

	// 读取flags参数
	//myname, _ := console.GetCmd().Flags().GetString("myname")
	//fmt.Println("myname=", myname)
	//
	//// 读取flags参数
	//environment, _ := console.GetCmd().Flags().GetString("environment")
	//fmt.Println("my-environment=", environment)

	//resDto := ListOrderDto{
	//	OrderId:  12,
	//	UserId:   23,
	//	JsonData: "",
	//	Address: map[string]any{
	//		"country": "中国",
	//		"city":    "深圳",
	//	},
	//}
	//fmt.Println(resDto)
	// resMap := map[string]any{}
	Order.returnJson(list)
}

// SaveOrder 保存订单信息
func (Order *OrderController) SaveOrder(ctx *gin.Context) {
	orderService := &LibraryOrder.OrderService{}
	orderId := orderService.SaveOrder()

	Order.returnJson(&map[string]int{
		"order_id": orderId,
	})
}
