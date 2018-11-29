// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import (
	"log"
	"net"
)

var BiSender *GameleySender

type GameleySender struct {
	Addr     string
	GameId   int
	RegionId int
	conn     *net.UDPConn
}

type BIModel interface {
	ToString() string
}

func NewGameleySender(addr string, gameId, regionId int) *GameleySender {
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatal("【BiSDK】", "配置Addr有误")
		return nil
	}
	udpConn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		log.Print("【BiSDK】", "连接日志服务失败")
		return nil
	}
	BiSender.Addr = addr
	BiSender.GameId = gameId
	BiSender.RegionId = regionId
	BiSender.conn = udpConn
	return BiSender
}

func (g *GameleySender) Send(biModel BIModel) {
	if g.conn == nil {
		log.Print("【BiSDK】", "连接日志服务异常")
		return
	}
	if _, err := g.conn.Write([]byte(biModel.ToString())); err != nil {
		log.Print("【BiSDK】", "发送日志失败", err)
		return
	}
	log.Print("【BiSDK】", "发送日志 ---> ", biModel.ToString())
}

func (g *GameleySender) Close() error {
	return g.conn.Close()
}
