// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import (
	"github.com/gameley-tc/bi-go"
	"strconv"
	"time"
)

type LogRole struct {
	LogPlat
	// 角色注册时间 在当前区服注册时间 跟reged无关
	regDt time.Time
	// 分区玩家id
	uid string
	// 分区id
	regionId int
	// 玩家等级
	level int
	// 注册过吗 是否在其他区服注册过
	reged int
	// vip等级
	vip int
	// 付过费吗
	payed int
}

func (l *LogRole) ToString() string {
	return bigo.BiJoin(l.uuid, l.uid, strconv.Itoa(l.platId), strconv.Itoa(l.regionId), strconv.Itoa(l.channelId), bigo.BiDateFormat(l.dt), bigo.BiDateFormat(l.regDt), strconv.Itoa(l.level), strconv.Itoa(l.reged), strconv.Itoa(l.vip), strconv.Itoa(l.payed))
}
