// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

type LogEquipPlayWay struct {
	LogReason
	// 装备类型
	equipType int
	//
	equipId int
	// 装备原来的等级
	equipOldLevel int
	// 装备新的等级
	equipNewLevel int
	// 本次变动的等级 非负数
	equipLevel int
	// 装备玩法类型
	playType int
}

