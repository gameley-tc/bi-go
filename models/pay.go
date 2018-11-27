// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import "github.com/gameley-tc/bi-go"

// 充值日志
type LogPay struct {
	LogReason
	// 充值类型 1代币 2月卡
	PayType bigo.LogEnumPay
	// 订单号
	orderNumber string
	// 充值代币金额
	// 建议实际充值单位*100,以避免打折等引起小数问题，默认为人民币分
	dPrice int64
	// 充值后存量
	// 充值后玩家拥有的代币（钻石）数
	dStoreNum int64
	// 充值后存量
	// 跟设置的充值金额单位对应，默认为人民币分
	totalNum int64
	// 首充标志位 只有当一个区服内的玩家第一次充值时发1，其他都发0
	// 0普通充值 1首次充值
	fPayFlag int
	// 充值渠道id 当登录渠道与充值渠道不同时需要发送，相同时发0
	payChannelId int
	// 内购项ID
	// 游戏用的统一计费点编号，代表一个计费点，一般充值回调会透传这个参数
	payId string
	// 内购项ID对应名称
	// 礼包名称及礼包中具体包含的各种物品名称都需要，中文，运营要能看懂
	payIdName string
	// 本次充值代币(钻石)数量
	dNum int
	// 充值前代币(钻石)存量
	// 充值前玩家拥有的代币（钻石）数
	dBeforeStoreNum int64
}

