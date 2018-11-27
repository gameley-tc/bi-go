// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

type LogEquip struct {
	LogReason
	// 装备类型
	EquipType int
	// 装备ID
	EquipId int
	// 装备数量
	Num int
}
