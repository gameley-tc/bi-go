// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

type LogAccount struct {
	*LogPlat
	// 游戏ID
	GameId int
}

func NewLogAccount(logPlat *LogPlat) *LogAccount {
	return &LogAccount{LogPlat: logPlat}
}
