// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogLoginU struct {
	*LogAccountRole
	*LogDevices
}

// 渠道账号登录日志必填字段
func NewLogLoginU(channelId int, uid string) *LogLoginU {
	return &LogLoginU{LogAccountRole: NewLogAccountRole(channelId, uid, uid), LogDevices: &LogDevices{}}
}

func (l *LogLoginU) ToString() string {
	return bigo.BiJoin("log_login_u", l.LogAccountRole.ToString(), l.LogDevices.ToString())
}
