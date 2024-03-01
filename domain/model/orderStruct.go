package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Order struct {
	OrderId  int    `json:"order_id"`
	UserId   int    `json:"user_id"`
	JsonData string `json:"json_data"`
}

func NewOrderModel() *Order {
	return &Order{}
}

// TableName 定义对应的数据表
func (order *Order) TableName() string {
	return "tbl_order"
}

func (order *Order) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AfterFind")
	return nil
}

func (order *Order) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("BeforeSave")
	return fmt.Errorf("save error")
}
