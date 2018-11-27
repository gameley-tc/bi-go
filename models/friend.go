// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import "github.com/gameley-tc/bi-go"

type LogFriends struct {
	LogReason
	// 增加或减少
	AddOrReduce bigo.LogEnumChange
	// 玩家原来的好友数量
	oldFriendsNum int
	// 玩家新的好友数量
	newFriendsNum int
	// 本次变动的好友数量
	friendsNum int
	// 好友类型
	friendType int
}

