// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import "time"

type LogBaseRole struct {
	LogBase
	// 角色注册时间 在当前区服注册时间 跟reged无关
	regDt time.Time
	// 分区玩家id
	uid string
	// 分区id
	regionId int
	// 玩家等级
	level int
	// 注册过吗 是否在其他区服注册过
	reged bool
	// vip等级
	vip int
	// 付过费吗
	payed int
}

