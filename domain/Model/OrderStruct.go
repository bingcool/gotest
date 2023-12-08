package Model

type Order struct {
	OrderId  int
	UserId   int
	JsonData string
}

// TableName 定义对应的数据表
func (Order) TableName() string {
	return "tbl_order"
}
