// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogItem struct {
	*LogReason
	// 增加或减少
	AddOrReduce int
	// 道具类型
	ItemType int
	// 道具ID
	ItemId int
	// 道具原来的数量
	OldNum int
	// 道具新的数量
	NewNum int
	// 本次变动的道具数量
	Num int
}

func (l *LogItem) ToString() string {
	return bigo.BiJoin("log_item", l.LogReason.ToString(), strconv.Itoa(l.AddOrReduce), strconv.Itoa(l.ItemType), strconv.Itoa(l.ItemId), strconv.Itoa(l.OldNum), strconv.Itoa(l.NewNum), strconv.Itoa(l.Num))
}

// 道具变动日志必填参数
// channelId 渠道id
// uid uid
// sequenceId 关联一次事件的唯一ID
// reason 一级原因
// subReason 二级原因
// itemType 道具类型
// itemId   道具id
// oldNum   道具原有数量
// newNum   道具新的数量
func NewLogItem(channelId int, uid string, sequenceId, reason, subReason string, itemType int, itemId int, oldNum int, newNum int) *LogItem {
	return &LogItem{LogReason: &LogReason{
		SequenceId: sequenceId,
		Reason:     reason,
		SubReason:  subReason,
		LogRole:    NewLogRole(channelId, uid),
	}, AddOrReduce: bigo.BiGetAddOrReduce(newNum, oldNum), ItemType: itemType, ItemId: itemId, OldNum: oldNum, NewNum: newNum, Num: bigo.BiAbs(newNum - oldNum)}
}
