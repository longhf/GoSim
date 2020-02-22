/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.07
 * Note:
 */

package proxy

import (
	"fmt"
	"github.com/longhf/gosim/common"
	"github.com/longhf/gosim/network/config"
	"github.com/longhf/gosim/network/types"
	"github.com/longhf/gosim/network/connection"
	hm "github.com/longhf/gosim/network/netutil/handlermanager"
	"net"
)


type ClientProxy struct {
	*connection.SimConnection
	clientId common.ClientId
	UdpAddr *net.UDPAddr
}


func NewClientProxy(conn net.Conn, clientId common.ClientId) *ClientProxy {
	return &ClientProxy{
		SimConnection: connection.NewSimConnection(conn),
		clientId:       clientId,
	}
}


func (cp *ClientProxy) Run(handlerManager *hm.HandlerManager) {
	cp.SetAutoFlush(config.CLIENT_PROXY_WRITE_FLUSH_INTERVAL)
	for {
		var msgType types.MsgType
		pkt, err := cp.Recv(&msgType)
		if err != nil {
			fmt.Println(err)
			break
		}
		data := pkt.ReadProtobufData()
		err = handlerManager.HandleMsg(msgType, data)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
