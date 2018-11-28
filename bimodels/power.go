// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

// 战力变动日志
type LogPower struct {
	*LogReason
	// 玩家原来的战力
	OldPower int64
	// 玩家新的战力
	NewPower int64
	// 本次变动的战力
	Power int64
}

func NewLogPower(logReason *LogReason, oldPower int64, newPower int64) *LogPower {
	return &LogPower{LogReason: logReason, OldPower: oldPower, NewPower: newPower, Power: bigo.BiAbsInt64(newPower - oldPower)}
}
