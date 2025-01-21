package entity

import (
	"fmt"
	"goTest/domain/dao/model"
	"gorm.io/gorm"
)

const TableNameTblOrder = "tbl_order"

// TblOrder 订单表
type TblOrder struct {
	model.TblOrder
}

// TableName TblOrder's table name
func (*TblOrder) TableName() string {
	return TableNameTblOrder
}

func NewOrderModel() *TblOrder {
	return &TblOrder{}
}

type Options func(order *TblOrder)

func WithAddress(address string) Options {
	return func(order *TblOrder) {
		order.Address = address
	}
}

func (user *TblOrder) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AfterFind")
	return
}

func (user *TblOrder) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("BeforeSave")
	return
}

func (user *TblOrder) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate")
	return
}
