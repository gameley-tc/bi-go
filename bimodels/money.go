// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogMoney struct {
	*LogReason
	// 原来的钱数
	OldMoney int64
	// 新的钱数
	NewMoney int64
	// 本次变动的钱数 非负数
	Money int64
	// 道具个数 本次购买道具、装备的个数 买货币时为0
	Num int
	// 增加或减少
	AddOrReduce int
	// 钱的类型
	// 注意：钻石货币类型必须为1，其他货币类型可自定义
	MoneyType int
	// 商品ID对应名称
	// 礼包名称及礼包中具体包含的各种物品名称都需要
	ItemIdName string
	// 商店编号 区分不同的商店
	ItemType int
	// 商品项ID 商店内唯一ID
	ItemId int
	// 是否属于商城购买 0否 1是
	Shop int
}

func (l *LogMoney) ToString() string {
	return bigo.BiJoin("log_money", l.LogReason.ToString(), strconv.FormatInt(l.OldMoney, 10), strconv.FormatInt(l.NewMoney, 10), strconv.FormatInt(l.Money, 10), strconv.Itoa(l.Num), strconv.Itoa(l.AddOrReduce), strconv.Itoa(l.MoneyType), l.ItemIdName, strconv.Itoa(l.ItemType), strconv.Itoa(l.ItemId), strconv.Itoa(l.Shop))
}

func NewLogMoney(logReason *LogReason, oldMoney int64, newMoney int64, num int, moneyType int, itemIdName string, itemType int, itemId int, shop int) *LogMoney {
	return &LogMoney{LogReason: logReason, OldMoney: oldMoney, NewMoney: newMoney, Num: num, AddOrReduce: bigo.BiGetAddOrReduceInt64(newMoney, oldMoney), MoneyType: moneyType, ItemIdName: itemIdName, ItemType: itemType, ItemId: itemId, Shop: shop, Money: bigo.BiAbsInt64(newMoney - oldMoney)}
}
