// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

type LogAccountRole struct {
	LogPlat
	// 游戏ID
	gameId int
	// 分区玩家id
	uid string
	// 分区id
	regionId int
	// 玩家等级
	level int
	// 注册过吗
	reged bool
}
