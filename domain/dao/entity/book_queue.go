package entity

import "goTest/domain/dao/model"

const TableNameBookQueue = "book_queue"

// BookQueue 预约队列任务表
type BookQueue struct {
	model.BookQueue
}

// TableName BookQueue's table name
func (*BookQueue) TableName() string {
	return TableNameBookQueue
}
