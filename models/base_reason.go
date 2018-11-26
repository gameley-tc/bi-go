// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import "github.com/gameley-tc/bi-go"

type LogBaseReason struct {
	LogBaseRole
	// 关联一次事件的唯一ID
	sequenceId string
	// 货币变动一级原因
	reason     string
	// 货币变动二级原因 没有填0
	subReason  string
	// 增加或减少
	addOrReduce bigo.TypesChange
}
