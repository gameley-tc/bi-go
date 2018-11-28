// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

// 任务日志表
type LogTask struct {
	*LogReason
	// 任务类型
	TaskType int
	// 任务ID
	TaskId int
}

func NewLogTask(logReason *LogReason, taskType int, taskId int) *LogTask {
	return &LogTask{LogReason: logReason, TaskType: taskType, TaskId: taskId}
}
