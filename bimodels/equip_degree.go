// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogEquipDegree struct {
	*LogReason
	// 装备类型
	EquipType int
	// 装备ID
	EquipId int
	// 装备原来的阶数
	EquipOldDegree int
	// 装备新的阶数
	EquipNewDegree int
	// 本次变动的阶数
	EquipDegree int
}

func NewLogEquipDegree(logReason *LogReason, equipType int, equipId int, equipOldDegree int, equipNewDegree int) *LogEquipDegree {
	return &LogEquipDegree{LogReason: logReason, EquipType: equipType, EquipId: equipId, EquipOldDegree: equipOldDegree, EquipNewDegree: equipNewDegree, EquipDegree: bigo.BiAbs(equipNewDegree - equipOldDegree)}
}
