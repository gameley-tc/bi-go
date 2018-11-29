// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogAction struct {
	*LogAccountRole
	*LogDevices
	// 日志类型
	// 1 进区服前行为 2 进区服后行为
	LogType bigo.LogEnumAction
	// 行为类型 游戏侧自定义
	ActionType int
	// 行为编号 游戏侧自定义
	ActionNumber int
	G1           int64
	G2           int64
	G3           int64
	G4           int64
	G5           int64
	G6           int64
	S1           string
	S2           string
	S3           string
	S4           string
	S5           string
	S6           string
}

func (l *LogAction) ToString() string {
	return bigo.BiJoin("log_action", l.LogAccountRole.ToString(), l.LogDevices.ToString(), strconv.Itoa(int(l.LogType)), strconv.Itoa(l.ActionType), strconv.Itoa(l.ActionNumber), strconv.FormatInt(l.G1, 10), strconv.FormatInt(l.G2, 10), strconv.FormatInt(l.G3, 10), strconv.FormatInt(l.G4, 10), strconv.FormatInt(l.G5, 10), strconv.FormatInt(l.G6, 10), l.S1, l.S2, l.S3, l.S4, l.S5, l.S6)
}

func NewLogAction(uuid, uid string, logType bigo.LogEnumAction, channelId, actionType, actionNumber int) *LogAction {
	return &LogAction{LogAccountRole: NewLogAccountRole(channelId, uid, uuid), LogDevices: &LogDevices{}, LogType: logType, ActionType: actionType, ActionNumber: actionNumber}
}
