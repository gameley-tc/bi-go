// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogGuildDonate struct {
	*LogReason
	// 	公会ID
	GuildId int
}

func (l *LogGuildDonate) ToString() string {
	return bigo.BiJoin("log_guild_donate", l.LogReason.ToString(), strconv.Itoa(l.GuildId))
}

func NewLogGuildDonate(logReason *LogReason, guildId int) *LogGuildDonate {
	return &LogGuildDonate{LogReason: logReason, GuildId: guildId}
}

