/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.10
 * Note:
 */

package server

import (
	pb "github.com/golang/protobuf/proto"
	"github.com/longhf/gosim/common"
	"github.com/longhf/gosim/network/proto"
	"github.com/longhf/gosim/network/config"
	"net"
)

type SimUDPConnection struct {
	conn net.UDPConn
	buffer []byte
	isClosed bool
}

func NewSimUDPConnection(conn net.UDPConn) *SimUDPConnection {
	return &SimUDPConnection{
		conn: conn,
		buffer: make([]byte, 1024 * 1024),
	}
}

func (uc *SimUDPConnection) Send(b []byte, addr *net.UDPAddr) (int, error) {
	return uc.conn.WriteToUDP(b, addr)
}

func (uc *SimUDPConnection) Recv(msgType *proto.MsgType) ([]byte, error) {
	cnt, _, err := uc.conn.ReadFromUDP(uc.buffer)
	if err != nil || cnt > 1048{  // TODO min max set
		return nil, err
	}
	*msgType = proto.MsgType(config.NETWORK_ENDIAN.Uint16(uc.buffer[:2]))
	data := uc.buffer[2:cnt]
	return data, nil
}


func (uc *SimUDPConnection) Close() error {
	uc.isClosed = true
	return uc.conn.Close()
}


func (uc *SimUDPConnection) IsClosed() bool {
	return uc.isClosed
}


// ------------------*- Sender -*----------------------
func (uc *SimUDPConnection) sendProtoData(proto interface{}, msgType proto.MsgType, dstAddr *net.UDPAddr) {
	data, err := pb.Marshal(proto.(pb.Message))
	if err != nil {
		return
	}
	pkt := make([]byte, 2 + len(data))
	util.NETWORK_ENDIAN.PutUint16(pkt[:2], uint16(msgType))
	copy(pkt[2:], data)
	uc.Send(pkt, dstAddr)
}


func (uc *SimUDPConnection) SendPosAndYaw(dstAddr *net.UDPAddr, clientId common.ClientId, entityId common.EntityId,
	x, y, z, yaw float32) {
	pAy := proto.PosAndYaw{
		ClientID: clientId.ToString(),
		EntityID: entityId.ToString(),
		PosX: x,
		PosY: y,
		PosZ: z,
		Yaw: yaw,
	}
	uc.sendProtoData(&pAy, proto.MT_POS_AND_YAW, dstAddr)
}
