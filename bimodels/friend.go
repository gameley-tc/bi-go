// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogFriends struct {
	*LogRole
	// 增加或减少
	AddOrReduce int
	// 玩家原来的好友数量
	OldFriendsNum int
	// 玩家新的好友数量
	NewFriendsNum int
	// 本次变动的好友数量
	FriendsNum int
	// 好友类型
	FriendType int
}

func (l *LogFriends) ToString(gameId string) string {
	return bigo.BiJoin("log_friends", l.LogRole.ToString(), strconv.Itoa(l.AddOrReduce), strconv.Itoa(l.OldFriendsNum), strconv.Itoa(l.NewFriendsNum), strconv.Itoa(l.FriendsNum), strconv.Itoa(l.FriendType))
}

func NewLogFriends(logRole *LogRole, oldFriendsNum int, newFriendsNum int, friendType int) *LogFriends {
	return &LogFriends{LogRole: logRole, AddOrReduce: bigo.BiGetAddOrReduce(newFriendsNum, oldFriendsNum), OldFriendsNum: oldFriendsNum, NewFriendsNum: newFriendsNum, FriendsNum: bigo.BiAbs(newFriendsNum - oldFriendsNum), FriendType: friendType}
}
