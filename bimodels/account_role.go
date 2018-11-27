// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

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

func NewLogAccountRole(logPlat *LogPlat, gameId int, uid string, regionId int, level int, reged int) *LogAccountRole {
	return &LogAccountRole{LogPlat: logPlat, GameId: gameId, Uid: uid, RegionId: regionId, Level: level, Reged: reged}
}
