// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import "github.com/gameley-tc/bi-go"

type LogGuildU struct {
	LogReason
	// 日志类型
	GuildUType bigo.LogEnumGuildU
	// 公会ID
	guildId int
}

