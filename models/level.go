// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

// 等级变动日志
type LogLevel struct {
	LogReason
	// 人物ID
	heroId int
	// 玩家新等级
	newLevel int
	// 本次变动的等级
	num int
}

