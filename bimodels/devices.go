// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bimodels

import "github.com/gameley-tc/bi-go"

type LogDevices struct {
	// 设备唯一id
	DeviceId string
	// ip
	Ip string
	// 客户端版本号
	ClientVersion string
	// 终端操作系统版本
	SystemSoftware string
	// 移动终端机型
	SystemHardware string
	// 运营商
	Telecom bigo.LogEnumTelecom
	// 网络类型
	Network bigo.LogEnumNetwork
	// cpu类型-频率-核数
	CpuHardware string
	// gpu类型-频率-核数
	GpuHardware string
	// 内存
	Memory int
	// 显示屏宽度
	ScreenWidth int
	// 显示屏高度
	ScreenHeight int
	// 像素密度
	Density float32
	// openGL render信息
	GlRender string
	// openGL 版本信息
	GlVersion string
	// 手机号
	PhoneNumber string
}
