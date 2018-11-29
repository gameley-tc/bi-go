// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"github.com/gameley-tc/bi-go"

	"strconv"
)

type LogAdVideo struct {
	*LogRole
	// 视频触发点
	Point int
	// 具体动作
	ActionType bigo.LogEnumAd
}

func NewLogAdVideo(logRole *LogRole, point int, actionType bigo.LogEnumAd) *LogAdVideo {
	return &LogAdVideo{LogRole: logRole, Point: point, ActionType: actionType}
}

func (l *LogAdVideo) ToString() string {
	return bigo.BiJoin("log_ad_video", l.LogRole.ToString(), strconv.Itoa(l.Point), strconv.Itoa(int(l.ActionType)))
}
