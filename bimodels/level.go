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

func (l *LogLevel) ToString(gameId string) string {
	return bigo.BiJoin("log_level", l.LogReason.ToString(), strconv.Itoa(l.HeroId), strconv.Itoa(l.NewLevel), strconv.Itoa(l.Num))
}

func NewLogLevel(logReason *LogReason, heroId int, newLevel int) *LogLevel {
	return &LogLevel{LogReason: logReason, HeroId: heroId, NewLevel: newLevel, Num: bigo.BiAbs(newLevel - logReason.Level)}
}
