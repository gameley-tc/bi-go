// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package main

import (
	"github.com/gameley-tc/bi-go"
	"github.com/gameley-tc/bi-go/bimodels"
)

func main() {
	bigo.InitBiSender(":8888", 12, 1010001)
	logAction := bimodels.NewLogAction("121212","wqeeqwe",bigo.LogEnumActionAfter, 2006603, 1212, 21121)
	logAdVideo := bimodels.NewLogAdVideo(2006603, "121212", 23, bigo.LogEnumAdMidway)
	logAppOnce := bimodels.NewLogAppOnce("121212", 2006603, 23)

	logBattle := bimodels.NewLogBattle(2006603, "121212", bigo.LogEnumGamePatternStart, 12, 1, 1212)

	logCdKey := bimodels.NewLogCdKey(2006603, "121212", 1212, "dsds", "sdsds", "sdsad")
	logCu := bimodels.NewLogCu(2006603, 1212, 1122)

	logEnergy := bimodels.NewLogEnergy(2006603, "121212", 233, 1313)
	logEquip := bimodels.NewLogEquip(2006603, "121212", "23", "sdds", "wqe", 2, 12, 2, 12)
	logEquipDegree := bimodels.NewLogEquipDegree(2006603, "121212", 12,12,12,21)
	logEquipLevel := bimodels.NewLogEquipLevel(2006603, "121212", 1,1,12,32)
	logEquipPlayWay := bimodels.NewLogEquipPlayWay(2006603, "121212", 12,12,12,32,1)

	logFriends := bimodels.NewLogFriends(2006603, "121212", 21,32,1)

	logGamePattern := bimodels.NewLogGamePattern(2006603, "121212",bigo.LogEnumGamePatternError, 12,12,12,12)
	logGuild := bimodels.NewLogGuild(2006603, "121212",bigo.LogEnumGuildCreate, 12)
	logGuildDonate := bimodels.NewLogGuildDonate(2006603, "121212",2)
	logGuildLevel := bimodels.NewLogGuildLevel(2006603, 21,212,1231)
	logGuildU := bimodels.NewLogGuildU(2006603, "121212",bigo.LogEnumGuildUJoin, 12)

	logItem := bimodels.NewLogItem(2006603, "121212","ds", "dd","sad",12,21,3,1)

	logLevel := bimodels.NewLogLevel(2006603, "121212",3,331)
	logLogin := bimodels.NewLogLogin(2006603, "121212",bigo.LogEnumUserLogin)
	logLoginU := bimodels.NewLogLoginU(2006603, "121212",)

	logMoney := bimodels.NewLogMoney(2006603, "121212","d","sd","s",23,13,32,1)

	logNewTask := bimodels.NewLogNewTask(2006603, "121212",21,12)

	logPay := bimodels.NewLogPay(2006603, "121212",bigo.LogEnumPayMoneyCard, "2312",123,23,12,1,2,2)
	logPower := bimodels.NewLogPower(2006603, "121212",2,12)

	logReg := bimodels.NewLogReg(2006603, "121212")
	logRegU := bimodels.NewLogRegU(2006603, "121212")

	logShareIn := bimodels.NewLogShareIn(2006603, "121212",1,1)
	logShareOut := bimodels.NewLogShareOut(2006603, "121212",1)

	logTask := bimodels.NewLogTask(2006603, "121212",1,1)

	bigo.BiSender.Send(logAction)
	bigo.BiSender.Send(logAdVideo)
	bigo.BiSender.Send(logAppOnce)
	bigo.BiSender.Send(logBattle)
	bigo.BiSender.Send(logCdKey)
	bigo.BiSender.Send(logCu)
	bigo.BiSender.Send(logEquip)
	bigo.BiSender.Send(logEquipDegree)
	bigo.BiSender.Send(logEquipLevel)
	bigo.BiSender.Send(logEquipPlayWay)
	bigo.BiSender.Send(logEnergy)
	bigo.BiSender.Send(logFriends)
	bigo.BiSender.Send(logGamePattern)
	bigo.BiSender.Send(logGamePattern)
	bigo.BiSender.Send(logGuildDonate)
	bigo.BiSender.Send(logGuildLevel)
	bigo.BiSender.Send(logGuildU)
	bigo.BiSender.Send(logGuild)
	bigo.BiSender.Send(logItem)
	bigo.BiSender.Send(logLevel)
	bigo.BiSender.Send(logLogin)
	bigo.BiSender.Send(logLoginU)
	bigo.BiSender.Send(logMoney)
	bigo.BiSender.Send(logNewTask)
	bigo.BiSender.Send(logPay)
	bigo.BiSender.Send(logPower)
	bigo.BiSender.Send(logReg)
	bigo.BiSender.Send(logRegU)
	bigo.BiSender.Send(logShareIn)
	bigo.BiSender.Send(logShareOut)
	bigo.BiSender.Send(logTask)

	bigo.BiSender.Close()
}
