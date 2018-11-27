// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

type LogItem struct {
	*LogReason
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
