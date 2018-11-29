// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

// 角色注册日志
type LogReg struct {
	*LogRole
	*LogDevices
}

func (l *LogReg) ToString() string {
	return bigo.BiJoin("log_reg", l.LogRole.ToStringReg(), l.LogDevices.ToString())
}

func NewLogReg(logRole *LogRole, logDevices *LogDevices) *LogReg {
	return &LogReg{LogRole: logRole, LogDevices: logDevices}
}

