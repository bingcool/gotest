package entity

import "goTest/domain/dao/model"

const TableNameCreditCard = "credit_cards"

// CreditCard mapped from table <credit_cards>
type CreditCard struct {
	model.CreditCard
}

// TableName CreditCard's table name
func (*CreditCard) TableName() string {
	return TableNameCreditCard
}
