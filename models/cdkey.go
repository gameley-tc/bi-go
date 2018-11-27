// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

type LogCdKey struct {
	LogRole
	// 活动ID
	ActionId int
	// 玩家输入的兑换码
	CdKey string
	// 奖品信息
	Prize string
	// 失败原因
	FailReason string
}
