package orderDto

// ListOrderReqDto 请求结构体
type ListOrderReqDto struct {
	OrderId  int `json:"order_id" form:"order_id" binding:"required,gt=0"`
	OrderId1 int `json:"order_id1" form:"order_id1"`
}

// ListOrderDto 响应结构体
type ListOrderDto struct {
	OrderId  int            `json:"order_id"`
	UserId   int            `json:"user_id"`
	JsonData string         `json:"json_data"`
	Address  map[string]any `json:"address"`
}
