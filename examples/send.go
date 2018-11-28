// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package main

import (
	"time"

	"github.com/gameley-tc/bi-go"
	"github.com/gameley-tc/bi-go/bimodels"
)

func main() {
	client := bigo.NewGameleySender(":8888", 12)
	logPlat := bimodels.NewLogPlat("1213313sdasd",2006603, time.Now())
	//logDevices := bimodels.NewLogDevices("sdasd322323", "1.1.1.", "sdad", "asda", "asdasda", 1, 1, "sadad","asdads",233,232,32,"sdsd","sdads","31312313")
	logDevices := &bimodels.LogDevices{}
	logRole := bimodels.NewLogRole(logPlat, time.Now(), "2323", 1010001, 12, 1, 1, 1)
	//logReason := bimodels.NewLogReason(logRole, "1212", "委屈翁群二", "0")
	//logAdVideo := bimodels.NewLogAdVideo(logRole, 5, bigo.LogEnumAdEnd)
	logReg := bimodels.NewLogReg(logRole, logDevices)
	//logEquip := bimodels.NewLogEquip(logReason, 1, 1, 1, 1)
	//client.Send(logAdVideo)
	client.Send(logReg)
	//client.Send(logEquip)
}
