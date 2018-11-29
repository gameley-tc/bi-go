// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogCdKey struct {
	*LogRole
	// 活动ID
	ActionId int
	// 玩家输入的兑换码
	CdKey string
	// 奖品信息
	Prize string
	// 失败原因
	FailReason string
}

func (l *LogCdKey) ToString() string {
	return bigo.BiJoin("log_cdkey", l.LogRole.ToString(), strconv.Itoa(l.ActionId), l.CdKey, l.Prize, l.FailReason)
}

func NewLogCdKey(channelId int, uid string, actionId int, cdKey string, prize string, failReason string) *LogCdKey {
	return &LogCdKey{LogRole: NewLogRole(channelId, uid), ActionId: actionId, CdKey: cdKey, Prize: prize, FailReason: failReason}
}
