// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

type LogGuildDonate struct {
	*LogReason
	// 	公会ID
	GuildId int
}

func NewLogGuildDonate(logReason *LogReason, guildId int) *LogGuildDonate {
	return &LogGuildDonate{LogReason: logReason, GuildId: guildId}
}
