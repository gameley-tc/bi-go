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

// money日志必填字段（非商城购买行为）
// channelId 渠道id
// uid  uid
// sequenceId 关联一次事件的唯一ID
// reason 一级原因
// subReason 二级原因
// oldMoney 代币老的数量
// newMoney 代币新的数量
// num  道具个数，本次购买道具、装备的个数 买货币时为0
// moneyType 代币类型
func NewLogMoney(channelId int, uid, sequenceId, reason, subReason string, oldMoney int64, newMoney int64, num int, moneyType int) *LogMoney {
	return &LogMoney{LogReason: &LogReason{
		SequenceId: sequenceId,
		Reason:     reason,
		SubReason:  subReason,
		LogRole:    NewLogRole(channelId, uid),
	}, OldMoney: oldMoney, NewMoney: newMoney, Num: num, AddOrReduce: bigo.BiGetAddOrReduceInt64(newMoney, oldMoney), MoneyType: moneyType, Money: bigo.BiAbsInt64(newMoney - oldMoney)}
}

// money日志必填字段（商城购买行为）
func NewLogMoneyStore(channelId int, uid, sequenceId, reason, subReason string, oldMoney int64, newMoney int64, num int, moneyType int, itemIdName string, itemType int, itemId int) *LogMoney {
	return &LogMoney{LogReason: &LogReason{
		SequenceId: sequenceId,
		Reason:     reason,
		SubReason:  subReason,
		LogRole:    NewLogRole(channelId, uid),
	}, OldMoney: oldMoney, NewMoney: newMoney, Num: num, AddOrReduce: bigo.BiGetAddOrReduceInt64(newMoney, oldMoney), MoneyType: moneyType, ItemIdName: itemIdName, ItemType: itemType, ItemId: itemId, Shop: 1, Money: bigo.BiAbsInt64(newMoney - oldMoney)}
}
