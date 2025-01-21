package entity

import "goTest/domain/dao/model"

const TableNameBank = "banks"

// Bank mapped from table <banks>
type Bank struct {
	model.Bank
}

// TableName Bank's table name
func (*Bank) TableName() string {
	return TableNameBank
}
