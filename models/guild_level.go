// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

type LogGuildLevel struct {
	LogGuildBase
	// 公会原来的等级
	guildOldLevel int
	// 公会新的等级
	guildNewLevel int
	// 本次变动的等级
	guildLevel int
}

