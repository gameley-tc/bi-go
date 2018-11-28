// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

// 新手任务日志表
type LogNewTask struct {
	*LogReason
	// 步骤ID
	TaskId int
	// 跟步骤ID对应的引导顺序
	OrderId int
}

func (l *LogNewTask) ToString(gameId string) string {
	return bigo.BiJoin("log_new_task", l.LogReason.ToString(), strconv.Itoa(l.TaskId), strconv.Itoa(l.OrderId))
}

func NewLogNewTask(logReason *LogReason, taskId int, orderId int) *LogNewTask {
	return &LogNewTask{LogReason: logReason, TaskId: taskId, OrderId: orderId}
}
