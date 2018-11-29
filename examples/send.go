// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package main

import (
	"github.com/gameley-tc/bi-go"
	"github.com/gameley-tc/bi-go/bimodels"
)

func main() {
	client := bigo.NewGameleySender(":8888", 12, 1010001)
	logPlat := bimodels.NewLogPlat(2006603)
	logDevices := &bimodels.LogDevices{}
	logRole := bimodels.NewLogRole(logPlat, "2323")
	logReason := bimodels.NewLogReason(logRole, "1212", "委屈翁群二", "0")
	logAdVideo := bimodels.NewLogAdVideo(logRole, 5, bigo.LogEnumAdEnd)
	logReg := bimodels.NewLogReg(logRole, logDevices)
	logEquip := bimodels.NewLogEquip(logReason, 1, 1, 1, 1)
	client.Send(logReg)
	client.Send(logEquip)
	client.Send(logAdVideo)
	client.Close()
}
