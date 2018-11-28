// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package main

import (
	"time"

	"github.com/gameley-tc/bi-go"
	"github.com/gameley-tc/bi-go/bimodels"
)

func main() {
	client := bigo.NewGameleySender(":8888", 12)
	logPlat := bimodels.NewLogPlat("1213313sdasd",1233123, time.Now())
	logRole := bimodels.NewLogRole(logPlat, time.Now(), "2323", 123123, 12, 1, 1, 1)
	logAdVideo := bimodels.NewLogAdVideo(logRole, 5, bigo.LogEnumAdEnd)
	client.Send(logAdVideo)
}
