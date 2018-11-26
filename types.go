// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

type TypesChange int
type TypesAction int
type TypesAd int
type TypesTelecom int
type TypesGamePattern int
type TypesGuild int
type TypesGuildU int
type TypesGot int
type TypesUserAction int
type TypesMoney int
type TypesNetwork int
type TypesOutput int
type TypesPay int
type TypesPlat int
type TypesStatus int

// 日志类型常量
const (
	TypesChangeReduce TypesChange = 0 // 减
	TypesChangeAdd    TypesChange = 1 // 加
	TypeChangeNone    TypesChange = 2 // 不变

	TypesActionBefore TypesAction = 1 // 进区服前行为
	TypesActionAfter  TypesAction = 2 // 进区服后行为

	TypesAdOpen   TypesAd = 0 // 主动触发打开视频窗口
	TypesAdClose  TypesAd = 1 // 看到视频窗口后主动关闭
	TypesAdClick  TypesAd = 2 // 点击视频播放视频
	TypesAdMidway TypesAd = 3 // 视频中途关闭
	TypesAdEnd    TypesAd = 4 // 观看结束
	TypesAdReward TypesAd = 5 // 获得奖励
	TypesAdNo     TypesAd = 6 // 没有可观看的广告

	TypesTelecomCMCC TypesTelecom = 1 // 中国移动
	TypesTelecomCUCC TypesTelecom = 2 // 中国联通
	TypesTelecomCTCC TypesTelecom = 3 // 中国电信
	TypesTelecomWIFI TypesTelecom = 4 // WIFI

	TypesGamePatternStart TypesGamePattern = 1 // 开始
	TypesGamePatternEnd   TypesGamePattern = 2 // 成功结束
	TypesGamePatternError TypesGamePattern = 3 // 失败结束

	TypesGuildCreate  TypesGuild = 1 // 公会创建
	TypesGuildDisband TypesGuild = 2 // 公会解散

	TypesGuildUCreate  TypesGuildU = 1 // 公会创建
	TypesGuildUDisband TypesGuildU = 2 // 公会解散

	TypesGotPay             TypesGot = 10001 // 充值
	TypesGotContinuousLogin TypesGot = 10002 // 连续登录活动

	TypesGotConsumeStore TypesGot = 1 // 商城购买

	TypesUserLogout TypesUserAction = 0 // 登录
	TypesUserLogin  TypesUserAction = 1 // 退出

	TypesMoneyIngot TypesMoney = 1 // 代币(钻石)
	TypesMoneyGold  TypesMoney = 2 // 游戏币(金币)

	TypesNetworkWIFI TypesNetwork = 1 // WIFI
	TypesNetwork4G   TypesNetwork = 2 // 4G
	TypesNetwork3G   TypesNetwork = 3 // 3G
	TypesNetwork2G   TypesNetwork = 4 // 2G

	TypesOutputMoney  TypesOutput = 1 // 游戏币
	TypesOutputEquip  TypesOutput = 2 // 装备
	TypesOutputItem   TypesOutput = 3 // 道具
	TypesOutputEnergy TypesOutput = 4 // 体力

	TypesPayIngot     TypesPay = 1 // 代币(钻石)
	TypesPayMoneyCard TypesPay = 2 // 月卡

	TypesPlatIOS      TypesPlat = 1 // 苹果
	TypesPlatIOSCrack TypesPlat = 2 // 苹果越狱
	TypesPlatAndroid  TypesPlat = 3 // 安卓

	TypesStatusStart TypesStatus = 1 // 开始
	TypesStatusEnd   TypesStatus = 1 // 结束(完成)
)
