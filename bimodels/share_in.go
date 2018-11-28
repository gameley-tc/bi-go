// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

// 分享入流水
type LogShareIn struct {
	*LogRole
	// 分享点
	Point int
	// 是否是新玩家
	NewPlayer int
}

func (l *LogShareIn) ToString(gameId string) string {
	return bigo.BiJoin("log_share_in", l.LogRole.ToString(), strconv.Itoa(l.Point), strconv.Itoa(l.NewPlayer))
}

func NewLogShareIn(logRole *LogRole, point int, newPlayer int) *LogShareIn {
	return &LogShareIn{LogRole: logRole, Point: point, NewPlayer: newPlayer}
}
