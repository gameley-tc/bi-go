// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import "github.com/gameley-tc/bi-go"

type LogGuild struct {
	LogReason
	// 日志类型
	GuildType bigo.LogEnumGuild
	// 公会ID
	guildId int
}

