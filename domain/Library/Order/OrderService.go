package Order

import (
	"goTest/domain/Factory"
	"goTest/domain/Model"
	"time"
)

type OrderService struct {
}

// GetOrderList 获取订单列表
func (OrderService *OrderService) GetOrderList(orderId int) []Model.Order {

	var orders []Model.Order
	db := Factory.GetDb()

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
	order := new(Model.Order)
	order.OrderId = int(seconds)
	order.UserId = 10000
	order.JsonData = "{}"

	db := Factory.GetDb()

	db.Select("order_id", "user_id", "json_data").Create(order)

	return order.OrderId
}
