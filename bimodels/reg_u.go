// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogRegU struct {
	*LogAccountRole
	*LogDevices
}

func NewLogRegU(channelId int, uuid string) *LogRegU {
	return &LogRegU{LogAccountRole: NewLogAccountRole(channelId, uuid, uuid), LogDevices: &LogDevices{}}
}

func (l *LogRegU) ToString() string {
	return bigo.BiJoin("log_reg_u", l.LogAccountRole.ToStringReg(), l.LogDevices.ToString())
}

