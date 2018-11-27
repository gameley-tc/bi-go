// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

type LogEnergy struct {
	*LogReason
	// 玩家原来的体力
	OldEnergy int
	// 玩家新的体力
	NewEnergy int
	// 本次变动的体力
	Energy int
}
