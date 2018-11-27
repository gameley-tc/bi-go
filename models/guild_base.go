// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import "time"

type LogGuildBase struct {
	// 平台ID
	platId int
	// 分区id
	regionId int
	// 时间
	dt time.Time
	// 关联一次事件的唯一ID
	sequenceId string
	// 货币变动一级原因
	reason string
	// 货币变动二级原因 没有填0
	subReason string
	// 公会ID
	guildId int
}

