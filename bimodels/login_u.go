// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

type LogLoginU struct {
	*LogAccountRole
	*LogDevices
}

func NewLogLoginU(logAccountRole *LogAccountRole, logDevices *LogDevices) *LogLoginU {
	return &LogLoginU{LogAccountRole: logAccountRole, LogDevices: logDevices}
}
