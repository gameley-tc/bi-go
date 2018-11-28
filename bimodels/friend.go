// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogFriends struct {
	*LogReason
	// 增加或减少
	AddOrReduce bigo.LogEnumChange
	// 玩家原来的好友数量
	OldFriendsNum int
	// 玩家新的好友数量
	NewFriendsNum int
	// 本次变动的好友数量
	FriendsNum int
	// 好友类型
	FriendType int
}

func NewLogFriends(logReason *LogReason, addOrReduce bigo.LogEnumChange, oldFriendsNum int, newFriendsNum int, friendType int) *LogFriends {
	return &LogFriends{LogReason: logReason, AddOrReduce: addOrReduce, OldFriendsNum: oldFriendsNum, NewFriendsNum: newFriendsNum, FriendType: friendType, FriendsNum: bigo.BiAbs(newFriendsNum - oldFriendsNum)}
}
