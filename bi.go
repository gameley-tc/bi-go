// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import (
	"net"
)

type GameleySender struct {
	Addr   string
	GameId int
	conn   *net.UDPConn
}

type BIModel interface {
	ToString() string
}

func NewGameleySender(addr string, gameId int) *GameleySender {
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		print("配置Addr有误")
		return nil
	}
	udpConn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		print("连接日志服务失败")
		return nil
	}
	return &GameleySender{Addr: addr, GameId: gameId, conn: udpConn}
}

func (g *GameleySender) Send(biModel BIModel) {
	if _, err := g.conn.Write([]byte(biModel.ToString())); err != nil {
		print("发送日志 ---->", biModel.ToString(), "失败")
	}
}
