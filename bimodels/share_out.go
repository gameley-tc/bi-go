// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

// 分享出流水
type LogShareOut struct {
	*LogRole
	// 分享点
	Point int
}

func NewLogShareOut(logRole *LogRole, point int) *LogShareOut {
	return &LogShareOut{LogRole: logRole, Point: point}
}
