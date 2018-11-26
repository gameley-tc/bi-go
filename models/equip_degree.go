// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

type  LogEquipDegree struct {
	LogBaseReason
	// 装备类型
	equipType int
	// 装备ID
	equipId int
	// 装备原来的阶数
	equipOldDegree int
	// 装备新的阶数
	equipNewDegree int
	// 本次变动的阶数
	equipDegree int
}

