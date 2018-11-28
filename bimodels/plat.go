// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"time"
)

type LogPlat struct {
	// 渠道账号唯一id
	Uuid string
	// 平台ID
	PlatId int
	// 渠道ID
	ChannelId int
	// 时间
	Dt time.Time
}

func NewLogPlat(uuid string, channelId int, dt time.Time) *LogPlat {
	return &LogPlat{Uuid: uuid, PlatId: channelId % 100, ChannelId: channelId, Dt: dt}
}
