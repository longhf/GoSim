/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.10
 * Note:
 */

package proxy

import (
	"fmt"
	"github.com/longhf/gosim/network/types"
	"github.com/longhf/gosim/network/connection"
	hm "github.com/longhf/gosim/network/netutil/handlermanager"
	"net"
)


type UdpProxy struct {
	*connection.SimUDPConnection
}


func NewUdpProxy(conn net.UDPConn) *UdpProxy {
	return &UdpProxy{
		SimUDPConnection:  connection.NewSimUDPConnection(conn),
	}
}


func (up *UdpProxy) Run(handlerManager *hm.HandlerManager) error {
	for{
		var msgType types.MsgType
		data, err := up.Recv(&msgType)
		if err != nil {
			fmt.Println(err)
			return err
		}
		go up.handleUDPPacket(data, msgType, handlerManager)  // TODO 若存在性能问题，加上goroutine pool
	}
}

func (up *UdpProxy) handleUDPPacket(data []byte, msgType types.MsgType, handlerManager *hm.HandlerManager) {
	err := handlerManager.HandleMsg(msgType, data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
