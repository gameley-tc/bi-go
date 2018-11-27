// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

// 战力变动日志
type LogPower struct {
	LogReason
	// 玩家原来的战力
	oldPower int64
	// 玩家新的战力
	newPower int64
	// 本次变动的战力
	power int64
}

