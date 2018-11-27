// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

type LogGuildLevel struct {
	*LogGuildBase
	// 公会原来的等级
	GuildOldLevel int
	// 公会新的等级
	GuildNewLevel int
	// 本次变动的等级
	GuildLevel int
}
