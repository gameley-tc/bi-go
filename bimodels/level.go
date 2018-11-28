// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

// 等级变动日志
type LogLevel struct {
	*LogReason
	// 人物ID
	HeroId int
	// 玩家新等级
	NewLevel int
	// 本次变动的等级
	Num int
}

func NewLogLevel(logReason *LogReason, heroId int, newLevel int, num int) *LogLevel {
	return &LogLevel{LogReason: logReason, HeroId: heroId, NewLevel: newLevel, Num: num}
}
