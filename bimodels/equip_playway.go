// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogEquipPlayWay struct {
	*LogReason
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

func (l *LogEquipPlayWay) ToString() string {
	return bigo.BiJoin("log_equip_playway", l.LogReason.ToString(), strconv.Itoa(l.EquipType), strconv.Itoa(l.EquipId), strconv.Itoa(l.EquipOldLevel), strconv.Itoa(l.EquipNewLevel),strconv.Itoa(l.EquipLevel), strconv.Itoa(l.PlayType))
}

func NewLogEquipPlayWay(channelId int, uid string, equipType int, equipId int, equipOldLevel int, equipNewLevel int, playType int) *LogEquipPlayWay {
	return &LogEquipPlayWay{LogReason: NewLogReason(channelId, uid), EquipType: equipType, EquipId: equipId, EquipOldLevel: equipOldLevel, EquipNewLevel: equipNewLevel, PlayType: playType, EquipLevel: bigo.BiAbs(equipNewLevel - equipOldLevel)}
}
