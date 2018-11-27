// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import "github.com/gameley-tc/bi-go"

type LogLogin struct {
	LogRole
	LogDevices
	// 登录类型
	LoginType bigo.LogEnumStatus
	// 玩家好友数量
	friendsNum int
	// 本次登录在线时间 退出时有,秒为单位
	online int
	// 玩家战力
	power int
	// 体力
	energy int
	// 公会ID 没有公会填0
	guildId string
	// 角色昵称 注意昵称中不能有|
	roleName string
}

