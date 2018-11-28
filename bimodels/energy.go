// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogEnergy struct {
	*LogReason
	// 增加或减少
	AddOrReduce int
	// 玩家原来的体力
	OldEnergy int
	// 玩家新的体力
	NewEnergy int
	// 本次变动的体力
	Energy int
}

func (l *LogEnergy) ToString(gameId string) string {
	return bigo.BiJoin("log_energy", l.LogReason.ToString(), strconv.Itoa(l.AddOrReduce), strconv.Itoa(l.OldEnergy), strconv.Itoa(l.NewEnergy), strconv.Itoa(l.Energy))
}

func NewLogEnergy(logReason *LogReason, oldEnergy int, newEnergy int) *LogEnergy {
	return &LogEnergy{LogReason: logReason, OldEnergy: oldEnergy, NewEnergy: newEnergy, Energy: bigo.BiAbs(newEnergy - oldEnergy), AddOrReduce: bigo.BiGetAddOrReduce(newEnergy, oldEnergy)}
}


