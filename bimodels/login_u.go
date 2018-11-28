// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogLoginU struct {
	*LogAccountRole
	*LogDevices
}

func NewLogLoginU(logAccountRole *LogAccountRole, logDevices *LogDevices) *LogLoginU {
	return &LogLoginU{LogAccountRole: logAccountRole, LogDevices: logDevices}
}

func (l *LogLoginU) ToString(gameId string) string{
	return bigo.BiJoin("log_login_u", l.LogAccountRole.ToString(gameId), l.LogDevices.ToString())
}
