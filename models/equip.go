// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

type LogEquip struct {
	LogReason
	// 装备类型
	equipType int
	// 装备ID
	equipId int
	// 装备数量
	num int
}

