// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

type LogEnergy struct {
	LogReason
	// 玩家原来的体力
	oldEnergy int
	// 玩家新的体力
	newEnergy int
	// 本次变动的体力
	energy int
}

