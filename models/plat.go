// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import "time"

type LogPlat struct {
	// 渠道账号唯一id
	uuid string
	// 平台ID
	platId int
	// 渠道ID
	channelId int
	// 时间
	dt time.Time
}