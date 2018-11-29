// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

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

func (l *LogLevel) ToString() string {
	return bigo.BiJoin("log_level", l.LogReason.ToString(), strconv.Itoa(l.HeroId), strconv.Itoa(l.NewLevel), strconv.Itoa(l.Num))
}

func NewLogLevel(channelId int, uid string, oldLevel int, newLevel int) *LogLevel {
	logReason := NewLogReason(channelId, uid)
	logReason.Level = oldLevel
	return &LogLevel{LogReason: NewLogReason(channelId, uid), NewLevel: newLevel, Num: bigo.BiAbs(newLevel - oldLevel)}
}
