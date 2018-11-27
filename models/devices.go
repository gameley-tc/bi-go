// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import "github.com/gameley-tc/bi-go"

type LogDevices struct {
	// 设备唯一id
	deviceId string
	// ip
	ip string
	// 客户端版本号
	clientVersion string
	// 终端操作系统版本
	systemSoftware string
	// 移动终端机型
	systemHardware string
	// 运营商
	Telecom bigo.LogEnumTelecom
	// 网络类型
	network bigo.LogEnumNetwork
	// cpu类型-频率-核数
	cpuHardware string
	// gpu类型-频率-核数
	gpuHardware string
	// 内存
	memory int
	// 显示屏宽度
	screenWidth int
	// 显示屏高度
	screenHeight int
	// 像素密度
	density float32
	// openGL render信息
	glRender string
	// openGL 版本信息
	glVersion string
	// 手机号
	phoneNumber string
}
