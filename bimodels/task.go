// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

// 任务日志表
type LogTask struct {
	*LogReason
	// 任务类型
	TaskType int
	// 任务ID
	TaskId int
}

func (l *LogTask) ToString(gameId string) string {
return bigo.BiJoin("log_task", l.LogReason.ToString(), strconv.Itoa(l.TaskType), strconv.Itoa(l.TaskId))
}

func NewLogTask(logReason *LogReason, taskType int, taskId int) *LogTask {
	return &LogTask{LogReason: logReason, TaskType: taskType, TaskId: taskId}
}
