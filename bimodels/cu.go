// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"
	"time"

	"github.com/gameley-tc/bi-go"
)

type LogCu struct {
	// 平台ID
	PlatId int
	// 区服ID
	RegionId int
	// 渠道ID
	ChannelId int
	// 时间
	Dt time.Time
	// 付费在线人数
	PayNum int
	// 非付费在线人数
	PayNotNum int
}

func (l *LogCu) ToString() string {
	return bigo.BiJoin("log_cu", strconv.Itoa(l.PlatId), strconv.Itoa(l.RegionId), strconv.Itoa(l.ChannelId), bigo.BiDateFormat(l.Dt), strconv.Itoa(l.PayNum), strconv.Itoa(l.PayNotNum))
}

func NewLogCu(channelId int, dt time.Time, payNum int, payNotNum int) *LogCu {
	return &LogCu{PlatId: channelId % 100, RegionId: bigo.BiSender.RegionId, ChannelId: channelId, Dt: dt, PayNum: payNum, PayNotNum: payNotNum}
}
