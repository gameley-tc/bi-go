// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"
	"time"

	"github.com/gameley-tc/bi-go"
)

type LogGuildBase struct {
	// 平台ID
	PlatId int
	// 渠道ID
	ChannelId int
	// 分区id
	RegionId int
	// 时间
	Dt time.Time
	// 关联一次事件的唯一ID
	SequenceId string
	// 货币变动一级原因
	Reason string
	// 货币变动二级原因 没有填0
	SubReason string
	// 公会ID
	GuildId int
}

func (l *LogGuildBase) ToString() string {
	return bigo.BiJoin(strconv.Itoa(l.PlatId), strconv.Itoa(l.RegionId), bigo.BiDateFormat(l.Dt), l.SequenceId, l.Reason, l.SubReason, strconv.Itoa(l.GuildId))
}

func NewLogGuildBase(channelId int, dt time.Time, sequenceId string, reason string, subReason string, guildId int) *LogGuildBase {
	return &LogGuildBase{PlatId: channelId % 100, ChannelId: channelId, RegionId: bigo.BiSender.RegionId, Dt: dt, SequenceId: sequenceId, Reason: reason, SubReason: subReason, GuildId: guildId}
}
