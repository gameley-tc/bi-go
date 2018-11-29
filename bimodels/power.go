// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

// 战力变动日志
type LogPower struct {
	*LogReason
	// 增加或减少
	AddOrReduce int
	// 玩家原来的战力
	OldPower int64
	// 玩家新的战力
	NewPower int64
	// 本次变动的战力
	Power int64
}

func (l *LogPower) ToString() string {
	return bigo.BiJoin("log_power", l.LogReason.ToString(), strconv.Itoa(l.AddOrReduce), strconv.FormatInt(l.OldPower, 10), strconv.FormatInt(l.NewPower, 10), strconv.FormatInt(l.Power, 10))
}

func NewLogPower(channelId int, uid string, oldPower int64, newPower int64) *LogPower {
	return &LogPower{LogReason: NewLogReason(channelId, uid), AddOrReduce: bigo.BiGetAddOrReduceInt64(newPower, oldPower), OldPower: oldPower, NewPower: newPower, Power: bigo.BiAbsInt64(newPower - oldPower)}
}
