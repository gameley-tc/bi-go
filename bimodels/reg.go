// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

// 角色注册日志
type LogReg struct {
	*LogRole
	*LogDevices
}

func NewLogReg(logRole *LogRole, logDevices *LogDevices) *LogReg {
	return &LogReg{LogRole: logRole, LogDevices: logDevices}
}
