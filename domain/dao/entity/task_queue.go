package entity

import "goTest/domain/dao/model"

const TableNameTaskQueue = "task_queue"

// TaskQueue 排队队列任务表
type TaskQueue struct {
	model.TaskQueue
}

// TableName TaskQueue's table name
func (*TaskQueue) TableName() string {
	return TableNameTaskQueue
}
