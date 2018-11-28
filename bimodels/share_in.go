// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

// 分享入流水
type LogShareIn struct {
	*LogRole
	// 分享点
	Point int
	// 是否是新玩家
	NewPlayer int
}

func NewLogShareIn(logRole *LogRole, point int, newPlayer int) *LogShareIn {
	return &LogShareIn{LogRole: logRole, Point: point, NewPlayer: newPlayer}
}
