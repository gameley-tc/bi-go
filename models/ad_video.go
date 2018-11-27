// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import (
	"github.com/gameley-tc/bi-go"
	"strconv"
)

type LogAdVideo struct {
	LogRole
	// 视频触发点
	Point int
	// 具体动作
	AdType bigo.LogEnumAd
}

func (l *LogAdVideo) ToString() string {
	return bigo.BiJoin("log_ad_video", l.LogRole.ToString(), strconv.Itoa(l.Point), strconv.Itoa(int(l.AdType)))
}
