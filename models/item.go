// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

type LogItem struct {
	LogReason
	// 道具类型
	itemType int
	// 道具ID
	itemId int
	// 道具原来的数量
	oldNum int
	// 道具新的数量
	newNum int
	// 本次变动的道具数量
	num int
}

