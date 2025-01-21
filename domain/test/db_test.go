package test

import (
	"goTest/domain/dao/entity"
	"goTest/domain/factory"
	"testing"
	"time"
)

func TestDb1(t *testing.T) {
	currentTime := time.Now()
	seconds := currentTime.Unix()
	order := entity.NewOrderModel()
	order.OrderID = int64(seconds)
	order.UserID = 10000
	order.ExpendData = "{}"
	order.JSONData = "{}"

	db := factory.GetDb()

	//db.Create(order)
	db.Save(order)
}
