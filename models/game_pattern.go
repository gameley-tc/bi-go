// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import "github.com/gameley-tc/bi-go"

type LogGamePattern struct {
	LogReason
	// 日志类型 1战斗开始 2战斗成功结束 3战斗失败结束
	GamePattern bigo.LogEnumGamePattern
	// 模式类型
	PatternType int
	// 模式ID
	PatternId int
	// 模式子ID 没有填0
	PatternSubId int
	// 本次闯关时间 单位秒 结束时有
	Times int
}
