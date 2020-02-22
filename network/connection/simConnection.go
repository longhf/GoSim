/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.07
 * Note: 基于连接的Connection及其Sender定义
 */

package util

import (
	pb "github.com/golang/protobuf/proto"
	"github.com/longhf/gosim/common"
	"github.com/longhf/gosim/network/proto"
	"net"
	"time"
)

type SimConnection struct{
	packetConn *PacketConnection
	autoFlushing bool
	isClosed bool
}

func NewSimConnection(conn net.Conn) *SimConnection {
	return &SimConnection{
		packetConn: NewPacketConnection(conn),
	}
}

func (sc *SimConnection) SendPacket(packet *Packet) error {
	return sc.packetConn.SendPacket(packet)
}


func (sc *SimConnection) Flush(reason string) error {
	// fmt.Println(reason)
	return sc.packetConn.Flush()
}

func (sc *SimConnection) SetAutoFlush(interval time.Duration) {
	if sc.autoFlushing {
		return
	}
	sc.autoFlushing = true
	go func() {
		for !sc.IsClosed() {
			time.Sleep(interval)
			err := sc.Flush("AutoFlush")
			if err != nil {
				break
			}
		}
	}()
}

func (sc *SimConnection) Recv(msgType *proto.MsgType) (*Packet, error) {
	pkt, err := sc.packetConn.RecvPacket()
	if err != nil {
		return nil, err
	}

	*msgType = proto.MsgType(pkt.ReadUint16())
	return pkt, nil
}


func (sc *SimConnection) Close() error {
	sc.isClosed = true
	return sc.packetConn.Close()
}


func (sc *SimConnection) IsClosed() bool {
	return sc.isClosed
}


// ------------------*- Sender -*----------------------
func (sc *SimConnection) sendProtoData(proto interface{}, msgType proto.MsgType) {
	data, err := pb.Marshal(proto.(pb.Message))
	if err != nil {
		return
	}
	pkt := NewPacket()
	pkt.AppendUint16(uint16(msgType))
	pkt.AppendUint32(uint32(len(data)))
	pkt.AppendProtobufData(data)
	sc.SendPacket(pkt)
}


func (sc *SimConnection) SendRegisterClient(addr string) {
	rc := proto.RegisterClient{
		UdpAddr: addr,
	}
	sc.sendProtoData(&rc, proto.MT_REGISTER_CLIENT)
}


func (sc *SimConnection) SendRegisterResult(clientId common.ClientId, result bool) {
	rr := proto.RegisterResult{
		Result:               result,
		ClientID:             clientId.ToString(),
	}
	sc.sendProtoData(&rr, proto.MT_REGISTER_RESULT)
}


func (sc *SimConnection) SendCreateEntity(clientId common.ClientId, entityId common.EntityId, entityType common.EntityType) {
	ce := proto.CreateEntity{
		ClientID: clientId.ToString(),
		EntityID: entityId.ToString(),
		EntityType: uint32(entityType),
	}
	sc.sendProtoData(&ce, proto.MT_CREATE_ENTITY)
}


func (sc *SimConnection) SendDestroyEntity(clientId common.ClientId, entityId common.EntityId) {
	de := proto.DestroyEntity{
		ClientID: clientId.ToString(),
		EntityID: entityId.ToString(),
	}
	sc.sendProtoData(&de, proto.MT_DESTROY_ENTITY)
}


func (sc *SimConnection) SendPosAndYaw(clientId common.ClientId, entityId common.EntityId, x, y, z, yaw float32) {
	pAy := proto.PosAndYaw{
		ClientID: clientId.ToString(),
		EntityID: entityId.ToString(),
		PosX: x,
		PosY: y,
		PosZ: z,
		Yaw: yaw,
	}
	sc.sendProtoData(&pAy, proto.MT_POS_AND_YAW)
}
