package entity

import "goTest/domain/dao/model"

const TableNameCustomer = "customers"

// Customer mapped from table <customers>
type Customer struct {
	model.Customer
}

// TableName Customer's table name
func (*Customer) TableName() string {
	return TableNameCustomer
}
