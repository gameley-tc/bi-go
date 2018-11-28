// Copyright 2018 The Gameley-TC Authors. All rights reserved.

package bigo

import (
	"log"
	"net"
	"strconv"
)

type GameleySender struct {
	Addr   string
	GameId int
	conn   *net.UDPConn
}

type BIModel interface {
	ToString(gameId string) string
}

func NewGameleySender(addr string, gameId int) *GameleySender {
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatal("【BiSDK】", "配置Addr有误")
		return nil
	}
	udpConn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		log.Print("连接日志服务失败")
		return nil
	}
	return &GameleySender{Addr: addr, GameId: gameId, conn: udpConn}
}

func (g *GameleySender) Send(biModel BIModel) {
	if _, err := g.conn.Write([]byte(biModel.ToString(strconv.Itoa(g.GameId)))); err != nil {
		log.Print("【BiSDK】", "发送日志失败", err)
		return
	}
	log.Print("【BiSDK】", "发送日志 ---> ", biModel.ToString(strconv.Itoa(g.GameId)), " 成功 ")
}

func (g *GameleySender) Close() error {
	return g.conn.Close()
}
