/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.16
 * Note:
 */

package proxy

import (
	"fmt"
	"github.com/longhf/gosim/network/config"
	"github.com/longhf/gosim/network/types"
	"github.com/longhf/gosim/network/connection"
	hm "github.com/longhf/gosim/network/netutil/handlermanager"
	"net"
)


type ServerProxy struct {
	*connection.SimConnection
}


func NewServerProxy(conn net.Conn) *ServerProxy {
	return &ServerProxy{
		SimConnection:     connection.NewSimConnection(conn),
	}
}


func (sp *ServerProxy) Run(handlerManager *hm.HandlerManager) error{
	sp.SetAutoFlush(config.SERVER_PROXY_WRITE_FLUSH_INTERVAL)
	for {
		var msgType types.MsgType
		pkt, err := sp.Recv(&msgType)
		if err != nil {
			fmt.Println(err)
			return err
		}
		data := pkt.ReadProtobufData()
		err = handlerManager.HandleMsg(msgType, data)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
}
