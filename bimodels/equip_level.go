// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogEquipLevel struct {
	*LogReason
	// 装备类型
	EquipType int
	// 装备ID
	EquipId int
	// 装备原来的等级
	EquipOldLevel int
	// 装备新的等级
	EquipNewLevel int
	// 本次变动的等级 非负数
	EquipLevel int
}

func NewLogEquipLevel(logReason *LogReason, equipType int, equipId int, equipOldLevel int, equipNewLevel int) *LogEquipLevel {
	return &LogEquipLevel{LogReason: logReason, EquipType: equipType, EquipId: equipId, EquipOldLevel: equipOldLevel, EquipNewLevel: equipNewLevel, EquipLevel: bigo.BiAbs(equipNewLevel - equipOldLevel)}
}
