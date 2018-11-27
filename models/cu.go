// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import "time"

type LogCu struct {
	// 平台ID
	PlatId int
	// 区服ID
	RegionId int
	// 渠道ID
	ChannelId int
	// 时间
	Dt time.Time
	// 付费在线人数
	PayNum int
	// 非付费在线人数
	PayNotNum int
}

