// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

type LogAccountRole struct {
	LogPlat
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
