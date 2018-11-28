// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

// 分享出流水
type LogShareOut struct {
	*LogRole
	// 分享点
	Point int
}

func (l *LogShareOut) ToString(gameId string) string {
	return bigo.BiJoin("log_share_out", l.LogRole.ToString(), strconv.Itoa(l.Point))
}

func NewLogShareOut(logRole *LogRole, point int) *LogShareOut {
	return &LogShareOut{LogRole: logRole, Point: point}
}
