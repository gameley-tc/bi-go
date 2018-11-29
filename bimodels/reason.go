// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogReason struct {
	*LogRole
	// 关联一次事件的唯一ID
	SequenceId string
	// 货币变动一级原因
	Reason string
	// 货币变动二级原因 没有填0
	SubReason string
}

func NewLogReason(channelId int, uid string) *LogReason {
	return &LogReason{LogRole: NewLogRole(channelId, uid)}
}

func (l *LogReason) ToString() string {
	return bigo.BiJoin(l.LogRole.ToString(), l.SequenceId, l.Reason, l.SubReason)
}
