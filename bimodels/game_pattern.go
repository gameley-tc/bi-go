// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogGamePattern struct {
	*LogReason
	// 日志类型 1战斗开始 2战斗成功结束 3战斗失败结束
	LogType bigo.LogEnumGamePattern
	// 模式类型
	PatternType int
	// 模式ID
	PatternId int
	// 模式子ID 没有填0
	PatternSubId int
	// 本次闯关时间 单位秒 结束时有
	Times int
}

func (l *LogGamePattern) ToString() string {
	return bigo.BiJoin("log_game_pattern", l.LogReason.ToString(), strconv.Itoa(int(l.LogType)), strconv.Itoa(l.PatternType), strconv.Itoa(l.PatternId), strconv.Itoa(l.PatternSubId), strconv.Itoa(l.Times))
}

func NewLogGamePattern(logReason *LogReason, logType bigo.LogEnumGamePattern, patternType int, patternId int, patternSubId int, times int) *LogGamePattern {
	return &LogGamePattern{LogReason: logReason, LogType: logType, PatternType: patternType, PatternId: patternId, PatternSubId: patternSubId, Times: times}
}

