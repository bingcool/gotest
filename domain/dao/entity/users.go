package entity

import "goTest/domain/dao/model"

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	model.User
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
