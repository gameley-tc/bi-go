// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogGuildLevel struct {
	*LogGuildBase
	// 公会原来的等级
	GuildOldLevel int
	// 公会新的等级
	GuildNewLevel int
	// 本次变动的等级
	GuildLevel int
}

func (l *LogGuildLevel) ToString() string {
	return bigo.BiJoin("log_guild_level", l.LogGuildBase.ToString(), strconv.Itoa(l.GuildOldLevel), strconv.Itoa(l.GuildNewLevel), strconv.Itoa(l.GuildLevel))
}

func NewLogGuildLevel(channelId, guildId, guildOldLevel, guildNewLevel int) *LogGuildLevel {
	return &LogGuildLevel{LogGuildBase: NewLogGuildBase(channelId, guildId), GuildOldLevel: guildOldLevel, GuildNewLevel: guildNewLevel, GuildLevel: bigo.BiAbs(guildNewLevel - guildOldLevel)}
}
