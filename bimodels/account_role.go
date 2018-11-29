// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogAccountRole struct {
	*LogPlat
	// 游戏ID
	GameId int
	// 分区玩家id
	Uid string
	// 分区id
	RegionId int
	// 玩家等级
	Level int
	// 注册过吗
	Reged int
}

func (l *LogAccountRole) ToString() string {
	return bigo.BiJoin(l.Uuid, l.Uid, strconv.Itoa(l.GameId), strconv.Itoa(l.PlatId), strconv.Itoa(l.RegionId), strconv.Itoa(l.ChannelId), bigo.BiDateFormat(l.Dt), strconv.Itoa(l.Level), strconv.Itoa(l.Reged))
}

func (l *LogAccountRole) ToStringU() string {
	return bigo.BiJoin(l.Uuid, strconv.Itoa(l.GameId), strconv.Itoa(l.PlatId), strconv.Itoa(l.ChannelId), bigo.BiDateFormat(l.Dt))
}

func NewLogAccountRole(channelId int, uid, uuid string) *LogAccountRole {
	return &LogAccountRole{LogPlat: NewLogPlat(channelId, uuid), GameId: bigo.BiSender.GameId, Uid: uid, RegionId: bigo.BiSender.RegionId}
}
