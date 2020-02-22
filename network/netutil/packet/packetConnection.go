/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.07
 * Note:
 */

package packet

import (
	"fmt"
	"github.com/longhf/gosim/common"
	"github.com/longhf/gosim/network/config"
	"net"
	"sync"
)


type PacketConnection struct{
	conn net.Conn
	pendingPackets     []*Packet
	pendingPacketsLock sync.Mutex

	payloadLenBuf         [config.SIZE_FIELD_SIZE]byte
	payloadLenBytesRecved int
	recvTotalPayloadLen   uint32
	recvedPayloadLen      uint32
	recvingPacket         *Packet
}


func NewPacketConnection(conn net.Conn) *PacketConnection {
	pc := &PacketConnection{
		conn:       conn,
	}

	return pc
}


func (pc *PacketConnection) SendPacket(packet *Packet) error {
	pc.pendingPacketsLock.Lock()
	pc.pendingPackets = append(pc.pendingPackets, packet)
	pc.pendingPacketsLock.Unlock()
	return nil
}


func (pc *PacketConnection) Flush() (err error) {
	pc.pendingPacketsLock.Lock()
	if len(pc.pendingPackets) == 0 {
		pc.pendingPacketsLock.Unlock()
		return
	}
	packets := make([]*Packet, 0, len(pc.pendingPackets))
	packets, pc.pendingPackets = pc.pendingPackets, packets
	pc.pendingPacketsLock.Unlock()

	if len(packets) == 1 {
		packet := packets[0]
		err = common.WriteAll(pc.conn, packet.data())
		return
	}

	for _, packet := range packets {
		common.WriteAll(pc.conn, packet.data())
	}
	return
}


func (pc *PacketConnection) RecvPacket() (*Packet, error) {
	if pc.payloadLenBytesRecved < config.SIZE_FIELD_SIZE {
		n, err := pc.conn.Read(pc.payloadLenBuf[pc.payloadLenBytesRecved:])
		pc.payloadLenBytesRecved += n
		if pc.payloadLenBytesRecved < config.SIZE_FIELD_SIZE {
			return nil, err
		}
		pc.recvTotalPayloadLen = NETWORK_ENDIAN.Uint32(pc.payloadLenBuf[:])

		if pc.recvTotalPayloadLen == 0 || pc.recvTotalPayloadLen > config.MAX_PAYLOAD_SIZE {
			err := fmt.Errorf("invalid payload length: %v", pc.recvTotalPayloadLen)
			pc.resetRecvStates()
			pc.Close()
			return nil, err
		}
		pc.recvedPayloadLen = 0
		pc.recvingPacket = NewPacket()
	}
	n, err := pc.conn.Read(pc.recvingPacket.bytes[config.SIZE_FIELD_SIZE + pc.recvedPayloadLen :
		config.SIZE_FIELD_SIZE + pc.recvTotalPayloadLen])
	pc.recvedPayloadLen += uint32(n)

	if pc.recvedPayloadLen == pc.recvTotalPayloadLen {
		packet := pc.recvingPacket
		pc.resetRecvStates()
		return packet, nil
	}
	return nil, err
}


func (pc *PacketConnection) resetRecvStates() {
	pc.payloadLenBytesRecved = 0
	pc.recvTotalPayloadLen = 0
	pc.recvedPayloadLen = 0
	pc.recvingPacket = nil
}


func (pc *PacketConnection) Close() error {
	return pc.conn.Close()
}
