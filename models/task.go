// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

// 任务日志表
type LogTask struct {
	LogReason
	// 任务类型
	taskType int
	// 任务ID
	taskId int
}