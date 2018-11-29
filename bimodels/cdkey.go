// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogCdKey struct {
	*LogRole
	// 活动ID
	ActionId string
	// 玩家输入的兑换码
	CdKey string
	// 奖品信息
	Prize string
	// 失败原因
	FailReason string
}

func (l *LogCdKey) ToString() string {
	return bigo.BiJoin("log_cdkey", l.LogRole.ToString(), l.ActionId, l.CdKey, l.Prize, l.FailReason)
}

func NewLogCdKey(logRole *LogRole, actionId string, cdKey string, prize string, failReason string) *LogCdKey {
	return &LogCdKey{LogRole: logRole, ActionId: actionId, CdKey: cdKey, Prize: prize, FailReason: failReason}
}
