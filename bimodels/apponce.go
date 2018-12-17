// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

// 打开APP只发一次日志必填字段
type LogAppOnce struct {
	*LogAccountRole
	*LogDevices
	// 行为编号 自定义
	ActionNumber int
}

func NewLogAppOnce(uuid string, channelId, actionNumber int) *LogAppOnce {
	return &LogAppOnce{LogAccountRole: NewLogAccountRole(channelId, uuid, uuid), LogDevices: &LogDevices{
		DeviceId:uuid,
	}, ActionNumber: actionNumber}
}

func (l *LogAppOnce) ToString() string {
	return bigo.BiJoin("log_app_once", l.LogAccountRole.ToString(), l.LogDevices.ToString(), strconv.Itoa(l.ActionNumber))
}
