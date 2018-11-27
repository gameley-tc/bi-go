// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

// 新手任务日志表
type LogNewTask struct {
	LogReason
	// 步骤ID
	taskId int
	// 跟步骤ID对应的引导顺序
	orderId int
}

