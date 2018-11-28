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

func BiAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func BiAbsInt64(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func BiGetAddOrReduce(new, old int) int {
	if (new - old) > 0 {
		return 1
	}
	if (new - old ) == 0 {
		return 2
	}
	return 0
}

