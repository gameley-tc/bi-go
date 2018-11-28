// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

// 新手任务日志表
type LogNewTask struct {
	*LogReason
	// 步骤ID
	TaskId int
	// 跟步骤ID对应的引导顺序
	OrderId int
}

func NewLogNewTask(logReason *LogReason, taskId int, orderId int) *LogNewTask {
	return &LogNewTask{LogReason: logReason, TaskId: taskId, OrderId: orderId}
}
