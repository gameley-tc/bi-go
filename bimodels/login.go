// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogLogin struct {
	*LogRole
	*LogDevices
	// 登录类型 0登出 1登入
	LoginType bigo.LogEnumStatus
	// 玩家好友数量
	FriendsNum int
	// 本次登录在线时间 退出时有,秒为单位
	Online int
	// 玩家战力
	Power int
	// 体力
	Energy int
	// 公会ID 没有公会填0
	GuildId string
	// 角色昵称 注意昵称中不能有|
	RoleName string
}

func NewLogLogin(logRole *LogRole, logDevices *LogDevices, loginType bigo.LogEnumStatus, friendsNum int, online int, power int, energy int, guildId string, roleName string) *LogLogin {
	return &LogLogin{LogRole: logRole, LogDevices: logDevices, LoginType: loginType, FriendsNum: friendsNum, Online: online, Power: power, Energy: energy, GuildId: guildId, RoleName: roleName}
}
