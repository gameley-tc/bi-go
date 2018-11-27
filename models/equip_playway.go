// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

type LogEquipPlayWay struct {
	LogReason
	// 装备类型
	EquipType int
	//
	EquipId int
	// 装备原来的等级
	EquipOldLevel int
	// 装备新的等级
	EquipNewLevel int
	// 本次变动的等级 非负数
	EquipLevel int
	// 装备玩法类型
	PlayType int
}
