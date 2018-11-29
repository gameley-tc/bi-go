// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogAppOnce struct {
	*LogAccountRole
	// 行为编号
	ActionNumber int
}

func NewLogAppOnce(logAccountRole *LogAccountRole, actionNumber int) *LogAppOnce {
	return &LogAppOnce{LogAccountRole: logAccountRole, ActionNumber: actionNumber}
}

func (l *LogAppOnce) ToString() string {
	return bigo.BiJoin("log_app_once", l.LogAccountRole.ToString(), strconv.Itoa(l.ActionNumber))
}
