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

func (l *LogAccountRole) ToString(gameId string) string {
	return bigo.BiJoin(l.Uuid, l.Uid, gameId, strconv.Itoa(l.PlatId), strconv.Itoa(l.RegionId), strconv.Itoa(l.ChannelId), bigo.BiDateFormat(l.Dt), strconv.Itoa(l.Level), strconv.Itoa(l.Reged))
}

func NewLogAccountRole(logPlat *LogPlat, uid string, regionId int, level int, reged int) *LogAccountRole {
	return &LogAccountRole{LogPlat: logPlat, Uid: uid, RegionId: regionId, Level: level, Reged: reged}
}
