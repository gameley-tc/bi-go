// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogGuild struct {
	*LogReason
	// 日志类型
	GuildType bigo.LogEnumGuild
	// 公会ID
	GuildId int
}

func NewLogGuild(logReason *LogReason, guildType bigo.LogEnumGuild, guildId int) *LogGuild {
	return &LogGuild{LogReason: logReason, GuildType: guildType, GuildId: guildId}
}
