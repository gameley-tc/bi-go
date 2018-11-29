// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogBattle struct {
	*LogReason
	// 日志类型 1战斗开始 2战斗成功结束 3战斗失败结束
	LogType bigo.LogEnumGamePattern
	// 关卡类型
	BattleType int
	// 关卡ID
	BattleId int
	// 本次闯关时间 单位秒 结束时有
	Times int
}

func (l *LogBattle) ToString() string {
	return bigo.BiJoin("log_battle", l.LogRole.ToString(), strconv.Itoa(int(l.LogType)), strconv.Itoa(l.BattleType), strconv.Itoa(l.BattleId), strconv.Itoa(l.Times))
}

func NewLogBattle(logReason *LogReason, logType bigo.LogEnumGamePattern, battleType int, battleId int, times int) *LogBattle {
	return &LogBattle{LogReason: logReason, LogType: logType, BattleType: battleType, BattleId: battleId, Times: times}
}


