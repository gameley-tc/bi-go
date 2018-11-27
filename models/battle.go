// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import "github.com/gameley-tc/bi-go"

type LogBattle struct {
	LogReason
	// 日志类型 1战斗开始 2战斗成功结束 3战斗失败结束
	LogType bigo.LogEnumGamePattern
	// 关卡类型
	BattleType int
	// 关卡ID
	BattleId int
	// 本次闯关时间 单位秒 结束时有
	Time int
}

