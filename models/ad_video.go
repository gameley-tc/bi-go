// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import (
	"github.com/gameley-tc/bi-go"
)

type LogAdVideo struct {
	LogRole
	// 视频触发点
	Point int
	// 具体动作
	AdType bigo.LogEnumAd
}
