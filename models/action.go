// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package models

import "github.com/gameley-tc/bi-go"

type LogAction struct {
	LogRole
	LogDevices
	LogType bigo.LogEnumAction
	actionType int
	actionNumber int
	g1 int64
	g2 int64
	g3 int64
	g4 int64
	g5 int64
	g6 int64

	s1 string
	s2 string
	s3 string
	s4 string
	s5 string
	s6 string
}

