// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import (
	"strings"
	"time"
)

func BiDateFormat(times time.Time) string {
	return times.Format("2006-01-02 15:04:05")
}

func BiJoin(args ...string) string {
	return strings.Join(args, "|")
}