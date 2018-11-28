// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogGuild struct {
	*LogReason
	// 日志类型
	LogType bigo.LogEnumGuild
	// 公会ID
	GuildId int
}

func (l *LogGuild) ToString(gameId string) string {
	return bigo.BiJoin("log_guild", l.LogReason.ToString(), strconv.Itoa(int(l.LogType)), strconv.Itoa(l.GuildId))
}

func NewLogGuild(logReason *LogReason, logType bigo.LogEnumGuild, guildId int) *LogGuild {
	return &LogGuild{LogReason: logReason, LogType: logType, GuildId: guildId}
}
