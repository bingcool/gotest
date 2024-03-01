package order

import (
	"fmt"
	"goTest/domain/factory"
	"goTest/domain/model"
	"gorm.io/gorm"
	"time"
)

type OrderService struct {
}

// NewOrderService 实例化对象
func NewOrderService() *OrderService {
	return &OrderService{}
}

// GetOrderList 获取订单列表
func (OrderService *OrderService) GetOrderList(orderId int) []model.Order {

	var orders []model.Order
	var orders1 []model.Order
	db := factory.GetDb()

	if orderId > 0 {
		db = db.Where("order_id = ?", orderId)
	}

	db = db.Where("user_id in ?", []int{101})

	db.Limit(10).Find(&orders)

	order := model.Order{}
	// var orders2 []model.Order
	// 原生sql打印出来
	sql1 := OrderService.ToSQL(func(db *gorm.DB) *gorm.DB {
		db = db.Table(order.TableName()).Where("user_id = ?", 102).Select("order_id", "user_id", "json_data").Find(orders1)
		return db
	})

	fmt.Println(sql1)
	//
	//db.Raw(sql1).Scan(&orders1)
	//
	//fmt.Println(orders1)

	//rows, _ := db.Table(order.TableName()).Where("user_id = ?", 101).Select("order_id", "user_id", "json_data").Rows()
	//defer func(rows *sql.Rows) {
	//	_ = rows.Close()
	//}(rows)
	//
	//for rows.Next() {
	//	orderNew := order
	//	rows.Scan(&orderNew.OrderId, &orderNew.UserId, &orderNew.JsonData)
	//	fmt.Println(orderNew)
	//}
	//
	//// 创建变量用于存储结果
	//var count int64
	//_ = db.Table(order.TableName()).Where("user_id = ?", 101).Count(&count)
	//
	//fmt.Println(count)
	//
	//_ = db.Table(order.TableName()).Where("user_id = ?", 101).Select("order_id", "user_id", "json_data").Scan(&orders1)

	fmt.Println(orders1)

	return orders
}

// SaveOrder 保存订单信息
func (OrderService *OrderService) SaveOrder() int {
	currentTime := time.Now()
	seconds := currentTime.Unix()
	order := model.NewOrderModel()
	order.OrderId = int(seconds)
	order.UserId = 10000
	order.JsonData = "{}"

	db := factory.GetDb()

	db.Create(order)

	return order.OrderId
}

// ToSQL for generate SQL string.
//
//	db.ToSQL(func(tx *gorm.DB) *gorm.DB {
//			return tx.Model(&User{}).Where(&User{Name: "foo", Age: 20})
//				.Limit(10).Offset(5)
//				.Order("name ASC")
//				.First(&User{})
//	})
func (OrderService *OrderService) ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string {
	tx := queryFn(factory.GetDb().Session(&gorm.Session{DryRun: true, SkipDefaultTransaction: true, NewDB: true}))
	stmt := tx.Statement

	return factory.GetDb().Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
}
