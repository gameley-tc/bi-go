// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

type LogAction struct {
	*LogAccountRole
	*LogDevices
	LogType      bigo.LogEnumAction
	ActionType   int
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

func (l *LogAction) ToString(gameId string) string {
	return bigo.BiJoin("log_action", l.LogAccountRole.ToString(gameId), strconv.Itoa(int(l.LogType)), strconv.Itoa(l.ActionType), strconv.Itoa(l.ActionNumber), strconv.FormatInt(l.G1, 10), strconv.FormatInt(l.G2, 10), strconv.FormatInt(l.G3, 10), strconv.FormatInt(l.G4, 10), strconv.FormatInt(l.G5, 10), strconv.FormatInt(l.G6, 10), l.S1, l.S2, l.S3, l.S4, l.S5, l.S6)
}

func NewLogAction(logAccountRole *LogAccountRole, logDevices *LogDevices, logType bigo.LogEnumAction, actionType int, actionNumber int) *LogAction {
	return &LogAction{LogAccountRole: logAccountRole, LogDevices: logDevices, LogType: logType, ActionType: actionType, ActionNumber: actionNumber}
}


