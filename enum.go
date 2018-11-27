// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

type LogEnumChange int
type LogEnumAction int
type LogEnumAd int
type LogEnumTelecom int
type LogEnumGamePattern int
type LogEnumGuild int
type LogEnumGuildU int
type LogEnumGot int
type LogEnumUserAction int
type LogEnumMoney int
type LogEnumNetwork int
type LogEnumOutput int
type LogEnumPay int
type LogEnumPlat int
type LogEnumStatus int

// 日志类型常量
const (
	LogEnumChangeReduce       LogEnumChange      = 0     // 减
	LogEnumChangeAdd          LogEnumChange      = 1     // 加
	LogEnumChangeNone         LogEnumChange      = 2     // 不变
	LogEnumActionBefore       LogEnumAction      = 1     // 进区服前行为
	LogEnumActionAfter        LogEnumAction      = 2     // 进区服后行为
	LogEnumAdOpen             LogEnumAd          = 0     // 主动触发打开视频窗口
	LogEnumAdClose            LogEnumAd          = 1     // 看到视频窗口后主动关闭
	LogEnumAdClick            LogEnumAd          = 2     // 点击视频播放视频
	LogEnumAdMidway           LogEnumAd          = 3     // 视频中途关闭
	LogEnumAdEnd              LogEnumAd          = 4     // 观看结束
	LogEnumAdReward           LogEnumAd          = 5     // 获得奖励
	LogEnumAdNo               LogEnumAd          = 6     // 没有可观看的广告
	LogEnumTelecomCMCC        LogEnumTelecom     = 1     // 中国移动
	LogEnumTelecomCUCC        LogEnumTelecom     = 2     // 中国联通
	LogEnumTelecomCTCC        LogEnumTelecom     = 3     // 中国电信
	LogEnumTelecomWIFI        LogEnumTelecom     = 4     // WIFI
	LogEnumGamePatternStart   LogEnumGamePattern = 1     // 开始
	LogEnumGamePatternEnd     LogEnumGamePattern = 2     // 成功结束
	LogEnumGamePatternError   LogEnumGamePattern = 3     // 失败结束
	LogEnumGuildCreate        LogEnumGuild       = 1     // 公会创建
	LogEnumGuildDisband       LogEnumGuild       = 2     // 公会解散
	LogEnumGuildUCreate       LogEnumGuildU      = 1     // 公会创建
	LogEnumGuildUDisband      LogEnumGuildU      = 2     // 公会解散
	LogEnumGotPay             LogEnumGot         = 10001 // 充值
	LogEnumGotContinuousLogin LogEnumGot         = 10002 // 连续登录活动
	LogEnumGotConsumeStore    LogEnumGot         = 1     // 商城购买
	LogEnumUserLogout         LogEnumUserAction  = 0     // 登录
	LogEnumUserLogin          LogEnumUserAction  = 1     // 退出
	LogEnumMoneyIngot         LogEnumMoney       = 1     // 代币(钻石)
	LogEnumMoneyGold          LogEnumMoney       = 2     // 游戏币(金币)
	LogEnumNetworkWIFI        LogEnumNetwork     = 1     // WIFI
	LogEnumNetwork4G          LogEnumNetwork     = 2     // 4G
	LogEnumNetwork3G          LogEnumNetwork     = 3     // 3G
	LogEnumNetwork2G          LogEnumNetwork     = 4     // 2G
	LogEnumOutputMoney        LogEnumOutput      = 1     // 游戏币
	LogEnumOutputEquip        LogEnumOutput      = 2     // 装备
	LogEnumOutputItem         LogEnumOutput      = 3     // 道具
	LogEnumOutputEnergy       LogEnumOutput      = 4     // 体力
	LogEnumPayIngot           LogEnumPay         = 1     // 代币(钻石)
	LogEnumPayMoneyCard       LogEnumPay         = 2     // 月卡
	LogEnumPlatIOS            LogEnumPlat        = 1     // 苹果
	LogEnumPlatIOSCrack       LogEnumPlat        = 2     // 苹果越狱
	LogEnumPlatAndroid        LogEnumPlat        = 3     // 安卓
	LogEnumStatusStart        LogEnumStatus      = 1     // 开始
	LogEnumStatusEnd          LogEnumStatus      = 1     // 结束(完成)
)
