// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

// 分享入流水
type LogShareIn struct {
	LogRole
	// 分享点
	point int
	// 是否是新玩家
	newPlayer int
}

