// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

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

func (l *LogLogin) ToString() string {
	return bigo.BiJoin("log_login", l.LogRole.ToString(), l.LogDevices.ToString(), strconv.Itoa(int(l.LoginType)), strconv.Itoa(l.FriendsNum), strconv.Itoa(l.Online), strconv.Itoa(l.Power), strconv.Itoa(l.Energy), l.GuildId, l.RoleName)
}

// 最基本的必填字段,如果还需要其它字段(online、guild_id、role_name)请自由组合
// channelId 渠道id
// loginType 登录类型
func NewLogLogin(channelId int, uid string, loginType bigo.LogEnumStatus) *LogLogin {
	return &LogLogin{LogRole: NewLogRole(channelId, uid), LogDevices: &LogDevices{}, LoginType: loginType}
}
