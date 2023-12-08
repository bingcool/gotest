package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goTest/domain/Go"
	librryOrder "goTest/domain/Library/Order"
	"log"
	"strconv"
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

type ListOrderDto struct {
	OrderId  int            `json:"order_id"`
	UserId   int            `json:"user_id"`
	JsonData string         `json:"json_data"`
	Address  map[string]any `json:"address"`
}

// ListOrder 订单列表
func (order *OrderController) ListOrder(c *gin.Context) {
	orderService := &librryOrder.OrderService{}
	result := orderService.GetOrderList(1685959471)
	var list []ListOrderDto
	i := 0
	for _, item := range result {
		listOrderDto := ListOrderDto{}
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

	m := map[string]any{"name": "bingcool"}
	//ch := make(chan any)
	Go.Run(func(ctx Go.Context) {
		fmt.Println(ctx.BucketMap["name"])
	}, Go.Context{
		BucketMap: m,
	})

	log.Println("This is a log message.")

	_, err := gin.DefaultWriter.Write([]byte("Request: " + c.Request.URL.Path + "\n"))
	if err != nil {
		fmt.Println(err.Error())
	}

	order.returnJson(list)
}

// SaveOrder 保存订单信息
func (order *OrderController) SaveOrder(c *gin.Context) {
	orderService := &librryOrder.OrderService{}
	orderId := orderService.SaveOrder()
	fmt.Println("订单ID-" + strconv.Itoa(orderId))
}
