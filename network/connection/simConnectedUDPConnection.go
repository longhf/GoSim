/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.10
 * Note:
 */

package connection

import (
	pb "github.com/golang/protobuf/proto"
	"github.com/longhf/gosim/common"
	"github.com/longhf/gosim/network/proto"
	"github.com/longhf/gosim/network/config"
	"net"
)

type SimConnectedUDPConnection struct {
	conn net.UDPConn
	buffer []byte
	isClosed bool
}

func NewSimConnectedUDPConnection(conn net.UDPConn) *SimConnectedUDPConnection {
	return &SimConnectedUDPConnection{
		conn: conn,
		buffer: make([]byte, 1024 * 1024),
	}
}

func (uc *SimConnectedUDPConnection) Send(b []byte) (int, error) {
	return uc.conn.Write(b)
}

func (uc *SimConnectedUDPConnection) Recv(msgType *proto.MsgType) ([]byte, error) {
	cnt, err := uc.conn.Read(uc.buffer)
	if err != nil || cnt > 1048{  // TODO min max set
		return nil, err
	}
	*msgType = proto.MsgType(config.NETWORK_ENDIAN.Uint16(uc.buffer[:2]))
	data := uc.buffer[2:cnt]
	return data, nil
}


func (uc *SimConnectedUDPConnection) Close() error {
	uc.isClosed = true
	return uc.conn.Close()
}


func (uc *SimConnectedUDPConnection) IsClosed() bool {
	return uc.isClosed
}


// ------------------*- Sender -*----------------------
func (uc *SimConnectedUDPConnection) sendProtoData(proto interface{}, msgType proto.MsgType) {
	data, err := pb.Marshal(proto.(pb.Message))
	if err != nil {
		return
	}
	pkt := make([]byte, 2 + len(data))
	config.NETWORK_ENDIAN.PutUint16(pkt[:2], uint16(msgType))
	copy(pkt[2:], data)
	uc.Send(pkt)
}

func (uc *SimConnectedUDPConnection) SendPosAndYaw(clientId common.ClientId, entityId common.EntityId,
	x, y, z, yaw float32) {
	pAy := proto.PosAndYaw{
		ClientID: clientId.ToString(),
		EntityID: entityId.ToString(),
		PosX: x,
		PosY: y,
		PosZ: z,
		Yaw: yaw,
	}
	uc.sendProtoData(&pAy, proto.MT_POS_AND_YAW)
}
