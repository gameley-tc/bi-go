// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogEquip struct {
	*LogReason
	// 增加或减少 0减少 1增加 2不变
	AddOrReduce int
	// 装备类型
	EquipType int
	// 装备ID
	EquipId int
	// 装备数量
	Num int
}

func NewLogEquip(logReason *LogReason, addOrReduce int, equipType int, equipId int, num int) *LogEquip {
	return &LogEquip{LogReason: logReason, AddOrReduce: addOrReduce, EquipType: equipType, EquipId: equipId, Num: num}
}

func (l *LogEquip) ToString() string {
	return bigo.BiJoin("log_equip", l.LogReason.ToString(), strconv.Itoa(l.AddOrReduce), strconv.Itoa(l.EquipType), strconv.Itoa(l.EquipId), strconv.Itoa(l.Num))
}




