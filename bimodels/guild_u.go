// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogGuildU struct {
	*LogReason
	// 日志类型
	LogType bigo.LogEnumGuildU
	// 公会ID
	GuildId int
}

func (l *LogGuildU) ToString() string {
	return bigo.BiJoin("log_guild_u", l.LogReason.ToString(), strconv.Itoa(int(l.LogType)), strconv.Itoa(l.GuildId))
}

func NewLogGuildU(channelId int, uid string, logType bigo.LogEnumGuildU, guildId int) *LogGuildU {
	return &LogGuildU{LogReason: NewLogReason(channelId, uid), LogType: logType, GuildId: guildId}
}
