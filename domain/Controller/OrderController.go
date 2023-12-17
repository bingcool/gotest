package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goTest/domain/Go"
	LibraryOrder "goTest/domain/Library/Order"
	"log"
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

// ListOrderReqDto 请求结构体
type ListOrderReqDto struct {
	OrderId  int `json:"order_id" form:"order_id"`
	OrderId1 int `json:"order_id1" form:"order_id1"`
}

// ListOrderDto 响应结构体
type ListOrderDto struct {
	OrderId  int            `json:"order_id"`
	UserId   int            `json:"user_id"`
	JsonData string         `json:"json_data"`
	Address  map[string]any `json:"address"`
}

// ListOrder 订单列表
func (Order *OrderController) ListOrder(c *gin.Context) {
	listOrderReqDto := ListOrderReqDto{}
	Order.bindToReqDtoStruct(&listOrderReqDto)

	//userId := Order.UserId
	fmt.Println("POST", listOrderReqDto.OrderId)
	fmt.Println("GET", listOrderReqDto.OrderId1)

	orderService := &LibraryOrder.OrderService{}
	result := orderService.GetOrderList(0)
	var list []ListOrderDto
	// list := make([]ListOrderDto, 0)
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

	// 读取flags参数
	//myname, _ := Console.GetCmd().Flags().GetString("myname")
	//fmt.Println("myname=", myname)
	//
	//// 读取flags参数
	//environment, _ := Console.GetCmd().Flags().GetString("environment")
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
func (Order *OrderController) SaveOrder(c *gin.Context) {
	orderService := &LibraryOrder.OrderService{}
	orderId := orderService.SaveOrder()

	Order.returnJson(&map[string]int{
		"order_id": orderId,
	})
}
