// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import "github.com/gameley-tc/bi-go"

type LogAction struct {
	LogRole
	LogDevices
	LogType      bigo.LogEnumAction
	ActionType   int
	ActionNumber int
	G1           int64
	G2           int64
	G3           int64
	G4           int64
	G5           int64
	G6           int64

	S1 string
	S2 string
	S3 string
	S4 string
	S5 string
	S6 string
}
