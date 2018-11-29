// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"
	"time"

	"github.com/gameley-tc/bi-go"
)

type LogRole struct {
	*LogPlat
	// 角色注册时间 在当前区服注册时间 跟reged无关
	RegDt time.Time
	// 分区玩家id
	Uid string
	// 分区id
	RegionId int
	// 玩家等级
	Level int
	// 注册过吗 是否在其他区服注册过
	Reged int
	// vip等级
	Vip int
	// 付过费吗
	Payed int
}

func NewLogRole(logPlat *LogPlat, uid string) *LogRole {
	logPlat.Uuid = uid
	return &LogRole{LogPlat: logPlat, Uid: uid, RegionId: bigo.BiSender.RegionId}
}

func (l *LogRole) ToString() string {
	return bigo.BiJoin(l.Uuid, l.Uid, strconv.Itoa(l.PlatId), strconv.Itoa(l.RegionId), strconv.Itoa(l.ChannelId), bigo.BiDateFormat(l.Dt), bigo.BiDateFormat(l.RegDt), strconv.Itoa(l.Level), strconv.Itoa(l.Reged), strconv.Itoa(l.Vip), strconv.Itoa(l.Payed))
}

func (l *LogRole) ToStringReg() string {
	return bigo.BiJoin(l.Uuid, l.Uid, strconv.Itoa(l.PlatId), strconv.Itoa(l.RegionId), strconv.Itoa(l.ChannelId), bigo.BiDateFormat(l.Dt), strconv.Itoa(l.Level), strconv.Itoa(l.Reged))
}
