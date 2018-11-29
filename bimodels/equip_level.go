// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

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

func (l *LogEquipLevel) ToString() string {
	return bigo.BiJoin("log_equip_level", l.LogReason.ToString(), strconv.Itoa(l.EquipType), strconv.Itoa(l.EquipId), strconv.Itoa(l.EquipOldLevel), strconv.Itoa(l.EquipNewLevel), strconv.Itoa(l.EquipLevel))
}

func NewLogEquipLevel(channelId int, uid string, equipType int, equipId int, equipOldLevel int, equipNewLevel int) *LogEquipLevel {
	return &LogEquipLevel{LogReason: NewLogReason(channelId, uid), EquipType: equipType, EquipId: equipId, EquipOldLevel: equipOldLevel, EquipNewLevel: equipNewLevel, EquipLevel: bigo.BiAbs(equipNewLevel - equipOldLevel)}
}
