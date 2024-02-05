package order

import (
	"goTest/domain/factory"
	"goTest/domain/model"
	"time"
)

type OrderService struct {
}

// NewOrderService 实例化对象
func NewOrderService() *OrderService {
	return &OrderService{}
}

// GetOrderList 获取订单列表
func (OrderService *OrderService) GetOrderList(orderId int) []model.Order {

	var orders []model.Order
	db := factory.GetDb()

	if orderId > 0 {
		db = db.Where("order_id > ?", orderId)
	}

	db.Limit(1).Find(&orders)
	orderList := make(map[string]any)
	orderList["order_id"] = 1234
	return orders
}

// SaveOrder 保存订单信息
func (OrderService *OrderService) SaveOrder() int {
	currentTime := time.Now()
	seconds := currentTime.Unix()
	order := new(model.Order)
	order.OrderId = int(seconds)
	order.UserId = 10000
	order.JsonData = "{}"

	db := factory.GetDb()

	db.Select("order_id", "user_id", "json_data").Create(order)

	return order.OrderId
}
