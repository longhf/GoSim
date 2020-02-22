/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.07
 * Note:
 */

package packet

import (
	"github.com/longhf/gosim/network/config"
	"unsafe"
)

var (
	packetEndian = NETWORK_ENDIAN
)

type Packet struct {
	readCursor uint32
	bytes []byte
}

func init(){
}

func allocPacket() *Packet{
	pkt := &Packet{
		readCursor:   0,
		bytes:        make([]byte, 1024 * 1024),
	}
	return pkt
}

func NewPacket() *Packet{
	return allocPacket()
}


func (p *Packet) data() []byte {
	return p.bytes[0 : config.PREPAYLOAD_SIZE+p.GetPayloadLen()]
}


func (p *Packet) GetPayloadLen() uint32 {
	return *(*uint32)(unsafe.Pointer(&p.bytes[0])) & config.PAYLOAD_LEN_MASK
}


func (p *Packet) AppendUint16(v uint16) {
	payloadEnd := config.PREPAYLOAD_SIZE + p.GetPayloadLen()
	packetEndian.PutUint16(p.bytes[payloadEnd:payloadEnd+2], v)
	*(*uint32)(unsafe.Pointer(&p.bytes[0])) += 2
}


func (p *Packet) AppendUint32(v uint32) {
	payloadEnd := config.PREPAYLOAD_SIZE + p.GetPayloadLen()
	packetEndian.PutUint32(p.bytes[payloadEnd:payloadEnd+4], v)
	*(*uint32)(unsafe.Pointer(&p.bytes[0])) += 4
}


func (p *Packet) AppendProtobufData(b []byte) {
	payloadEnd := config.PREPAYLOAD_SIZE + p.GetPayloadLen()
	dataLen := uint32(len(b))
	copy(p.bytes[payloadEnd: payloadEnd + dataLen], b)
	*(*uint32)(unsafe.Pointer(&p.bytes[0])) += dataLen
}

func (p *Packet) ReadUint16() (v uint16) {
	pos := p.readCursor + config.PREPAYLOAD_SIZE
	v = packetEndian.Uint16(p.bytes[pos : pos+2])
	p.readCursor += 2
	return
}


func (p *Packet) ReadUint32() (v uint32) {
	pos := p.readCursor + config.PREPAYLOAD_SIZE
	v = packetEndian.Uint32(p.bytes[pos : pos+4])
	p.readCursor += 4
	return
}


func (p *Packet) ReadProtobufData() (v []byte) {
	size := p.ReadUint32()
	pos := p.readCursor + config.PREPAYLOAD_SIZE
	v = p.bytes[pos : pos + size]
	p.readCursor += size + 4
	return
}
