package model

import (
	"gorm.io/gorm"
)

func (order *TblOrder) AfterFind(tx *gorm.DB) (err error) {
	return nil
}

func (order *TblOrder) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}

// gentool -dsn "user:pwd@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local" -tables "orders,doctor"
