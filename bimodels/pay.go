// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import (
	"strconv"

	"github.com/gameley-tc/bi-go"
)

// 充值日志
type LogPay struct {
	*LogRole
	// 充值类型 1代币 2月卡
	PayType bigo.LogEnumPay
	// 订单号
	OrderNumber string
	// 充值代币金额
	// 建议实际充值单位*100,以避免打折等引起小数问题，默认为人民币分
	DPrice int64
	// 充值后存量
	// 充值后玩家拥有的代币（钻石）数
	DStoreNum int64
	// 截止该笔累计充值金额
	// 跟设置的充值金额单位对应，默认为人民币分
	TotalNum int64
	// 首充标志位 只有当一个区服内的玩家第一次充值时发1，其他都发0
	// 0普通充值 1首次充值
	FPayFlag int
	// 充值渠道id 当登录渠道与充值渠道不同时需要发送，相同时发0
	PayChannelId int
	// 内购项ID
	// 游戏用的统一计费点编号，代表一个计费点，一般充值回调会透传这个参数
	PayId string
	// 内购项ID对应名称
	// 礼包名称及礼包中具体包含的各种物品名称都需要，中文，运营要能看懂
	PayIdName string
	// 本次充值代币(钻石)数量
	DNum int
	// 充值前代币(钻石)存量
	// 充值前玩家拥有的代币（钻石）数
	DBeforeStoreNum int64
}

func (l *LogPay) ToString() string {
	return bigo.BiJoin("log_pay", l.LogRole.ToString(), strconv.Itoa(int(l.PayType)), l.OrderNumber, strconv.FormatInt(l.DPrice, 10), strconv.FormatInt(l.DStoreNum, 10), strconv.FormatInt(l.TotalNum, 10), strconv.Itoa(l.FPayFlag), strconv.Itoa(l.PayChannelId), l.PayId, l.PayIdName, strconv.Itoa(l.DNum), strconv.FormatInt(l.DBeforeStoreNum, 10))
}

// 充值日志必填字段
// channelId 渠道id
// uid  uid
// payTypeEnum  充值类型，详情请看枚举
// orderNumber  订单号
// dPrice     充值代币金额，建议实际充值单位*100,以避免打折等引起小数问题，默认为人民币分
// dStoreNum  充值后存量，充值后玩家拥有的代币（钻石）数
// totalNum   首充标志位 只有当一个区服内的玩家第一次充值时发1，其他都发0
// fPayFlag   充值渠道id 当登录渠道与充值渠道不同时需要发送，相同时发0
// dNum       本次充值代币(钻石)数量
// dBeforeStoreNum  充值前代币(钻石)存量
func NewLogPay(channelId int, uid string, payType bigo.LogEnumPay, orderNumber string, dPrice int64, dStoreNum int64, totalNum int64, fPayFlag int, dNum int, dBeforeStoreNum int64) *LogPay {
	return &LogPay{LogRole: NewLogRole(channelId, uid), PayType: payType, OrderNumber: orderNumber, DPrice: dPrice, DStoreNum: dStoreNum, TotalNum: totalNum, FPayFlag: fPayFlag, DNum: dNum, DBeforeStoreNum: dBeforeStoreNum}
}


