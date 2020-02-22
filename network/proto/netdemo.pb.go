// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: netdemo.proto

package proto

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterClient struct {
	UdpAddr              string   `protobuf:"bytes,1,opt,name=udpAddr,proto3" json:"udpAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterClient) Reset()         { *m = RegisterClient{} }
func (m *RegisterClient) String() string { return proto.CompactTextString(m) }
func (*RegisterClient) ProtoMessage()    {}
func (*RegisterClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fee4afce7facf50, []int{0}
}
func (m *RegisterClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterClient.Merge(m, src)
}
func (m *RegisterClient) XXX_Size() int {
	return m.Size()
}
func (m *RegisterClient) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterClient.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterClient proto.InternalMessageInfo

func (m *RegisterClient) GetUdpAddr() string {
	if m != nil {
		return m.UdpAddr
	}
	return ""
}

type RegisterResult struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	ClientID             string   `protobuf:"bytes,2,opt,name=clientID,proto3" json:"clientID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResult) Reset()         { *m = RegisterResult{} }
func (m *RegisterResult) String() string { return proto.CompactTextString(m) }
func (*RegisterResult) ProtoMessage()    {}
func (*RegisterResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fee4afce7facf50, []int{1}
}
func (m *RegisterResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResult.Merge(m, src)
}
func (m *RegisterResult) XXX_Size() int {
	return m.Size()
}
func (m *RegisterResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResult.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResult proto.InternalMessageInfo

func (m *RegisterResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *RegisterResult) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

type CreateEntity struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	EntityID             string   `protobuf:"bytes,2,opt,name=entityID,proto3" json:"entityID,omitempty"`
	EntityType           uint32   `protobuf:"varint,3,opt,name=entityType,proto3" json:"entityType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntity) Reset()         { *m = CreateEntity{} }
func (m *CreateEntity) String() string { return proto.CompactTextString(m) }
func (*CreateEntity) ProtoMessage()    {}
func (*CreateEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fee4afce7facf50, []int{2}
}
func (m *CreateEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntity.Merge(m, src)
}
func (m *CreateEntity) XXX_Size() int {
	return m.Size()
}
func (m *CreateEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntity.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntity proto.InternalMessageInfo

func (m *CreateEntity) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *CreateEntity) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *CreateEntity) GetEntityType() uint32 {
	if m != nil {
		return m.EntityType
	}
	return 0
}

type DestroyEntity struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	EntityID             string   `protobuf:"bytes,2,opt,name=entityID,proto3" json:"entityID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyEntity) Reset()         { *m = DestroyEntity{} }
func (m *DestroyEntity) String() string { return proto.CompactTextString(m) }
func (*DestroyEntity) ProtoMessage()    {}
func (*DestroyEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fee4afce7facf50, []int{3}
}
func (m *DestroyEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DestroyEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DestroyEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DestroyEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyEntity.Merge(m, src)
}
func (m *DestroyEntity) XXX_Size() int {
	return m.Size()
}
func (m *DestroyEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyEntity.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyEntity proto.InternalMessageInfo

func (m *DestroyEntity) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *DestroyEntity) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

type PosAndYaw struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	EntityID             string   `protobuf:"bytes,2,opt,name=entityID,proto3" json:"entityID,omitempty"`
	PosX                 float32  `protobuf:"fixed32,3,opt,name=pos_x,json=posX,proto3" json:"pos_x,omitempty"`
	PosY                 float32  `protobuf:"fixed32,4,opt,name=pos_y,json=posY,proto3" json:"pos_y,omitempty"`
	PosZ                 float32  `protobuf:"fixed32,5,opt,name=pos_z,json=posZ,proto3" json:"pos_z,omitempty"`
	Yaw                  float32  `protobuf:"fixed32,6,opt,name=yaw,proto3" json:"yaw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PosAndYaw) Reset()         { *m = PosAndYaw{} }
func (m *PosAndYaw) String() string { return proto.CompactTextString(m) }
func (*PosAndYaw) ProtoMessage()    {}
func (*PosAndYaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fee4afce7facf50, []int{4}
}
func (m *PosAndYaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PosAndYaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PosAndYaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PosAndYaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PosAndYaw.Merge(m, src)
}
func (m *PosAndYaw) XXX_Size() int {
	return m.Size()
}
func (m *PosAndYaw) XXX_DiscardUnknown() {
	xxx_messageInfo_PosAndYaw.DiscardUnknown(m)
}

var xxx_messageInfo_PosAndYaw proto.InternalMessageInfo

func (m *PosAndYaw) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *PosAndYaw) GetEntityID() string {
	if m != nil {
		return m.EntityID
	}
	return ""
}

func (m *PosAndYaw) GetPosX() float32 {
	if m != nil {
		return m.PosX
	}
	return 0
}

func (m *PosAndYaw) GetPosY() float32 {
	if m != nil {
		return m.PosY
	}
	return 0
}

func (m *PosAndYaw) GetPosZ() float32 {
	if m != nil {
		return m.PosZ
	}
	return 0
}

func (m *PosAndYaw) GetYaw() float32 {
	if m != nil {
		return m.Yaw
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterClient)(nil), "proto.RegisterClient")
	proto.RegisterType((*RegisterResult)(nil), "proto.RegisterResult")
	proto.RegisterType((*CreateEntity)(nil), "proto.CreateEntity")
	proto.RegisterType((*DestroyEntity)(nil), "proto.DestroyEntity")
	proto.RegisterType((*PosAndYaw)(nil), "proto.PosAndYaw")
}

func init() { proto.RegisterFile("netdemo.proto", fileDescriptor_3fee4afce7facf50) }

var fileDescriptor_3fee4afce7facf50 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4b, 0x2d, 0x49,
	0x49, 0xcd, 0xcd, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x5a, 0x5c,
	0x7c, 0x41, 0xa9, 0xe9, 0x99, 0xc5, 0x25, 0xa9, 0x45, 0xce, 0x39, 0x99, 0xa9, 0x79, 0x25, 0x42,
	0x12, 0x5c, 0xec, 0xa5, 0x29, 0x05, 0x8e, 0x29, 0x29, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x30, 0xae, 0x92, 0x0b, 0x42, 0x6d, 0x50, 0x6a, 0x71, 0x69, 0x4e, 0x89, 0x90, 0x18, 0x17,
	0x5b, 0x11, 0x98, 0x05, 0x56, 0xca, 0x11, 0x04, 0xe5, 0x09, 0x49, 0x71, 0x71, 0x24, 0x83, 0x4d,
	0xf3, 0x74, 0x91, 0x60, 0x02, 0x1b, 0x02, 0xe7, 0x2b, 0xa5, 0x71, 0xf1, 0x38, 0x17, 0xa5, 0x26,
	0x96, 0xa4, 0xba, 0xe6, 0x95, 0x64, 0x96, 0x54, 0xa2, 0xa8, 0x65, 0x44, 0x55, 0x0b, 0x92, 0x4b,
	0x05, 0xab, 0x42, 0x98, 0x03, 0xe3, 0x0b, 0xc9, 0x71, 0x71, 0x41, 0xd8, 0x21, 0x95, 0x05, 0xa9,
	0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x48, 0x22, 0x4a, 0xee, 0x5c, 0xbc, 0x2e, 0xa9, 0xc5,
	0x25, 0x45, 0xf9, 0x95, 0x94, 0x59, 0xa4, 0x34, 0x85, 0x91, 0x8b, 0x33, 0x20, 0xbf, 0xd8, 0x31,
	0x2f, 0x25, 0x32, 0xb1, 0x9c, 0x6c, 0xe7, 0x0a, 0x73, 0xb1, 0x16, 0xe4, 0x17, 0xc7, 0x57, 0x80,
	0x5d, 0xca, 0x14, 0xc4, 0x52, 0x90, 0x5f, 0x1c, 0x01, 0x13, 0xac, 0x94, 0x60, 0x81, 0x0b, 0x46,
	0xc2, 0x04, 0xab, 0x24, 0x58, 0xe1, 0x82, 0x51, 0x42, 0x02, 0x5c, 0xcc, 0x95, 0x89, 0xe5, 0x12,
	0x6c, 0x60, 0x21, 0x10, 0xd3, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x1c, 0xa5, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x35, 0x40, 0xab, 0xf2, 0xea, 0x01, 0x00, 0x00,
}

func (m *RegisterClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.UdpAddr) > 0 {
		i -= len(m.UdpAddr)
		copy(dAtA[i:], m.UdpAddr)
		i = encodeVarintNetdemo(dAtA, i, uint64(len(m.UdpAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintNetdemo(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0x12
	}
	if m.Result {
		i--
		if m.Result {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CreateEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateEntity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateEntity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EntityType != 0 {
		i = encodeVarintNetdemo(dAtA, i, uint64(m.EntityType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.EntityID) > 0 {
		i -= len(m.EntityID)
		copy(dAtA[i:], m.EntityID)
		i = encodeVarintNetdemo(dAtA, i, uint64(len(m.EntityID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintNetdemo(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DestroyEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DestroyEntity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DestroyEntity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.EntityID) > 0 {
		i -= len(m.EntityID)
		copy(dAtA[i:], m.EntityID)
		i = encodeVarintNetdemo(dAtA, i, uint64(len(m.EntityID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintNetdemo(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PosAndYaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PosAndYaw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PosAndYaw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Yaw != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Yaw))))
		i--
		dAtA[i] = 0x35
	}
	if m.PosZ != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.PosZ))))
		i--
		dAtA[i] = 0x2d
	}
	if m.PosY != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.PosY))))
		i--
		dAtA[i] = 0x25
	}
	if m.PosX != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.PosX))))
		i--
		dAtA[i] = 0x1d
	}
	if len(m.EntityID) > 0 {
		i -= len(m.EntityID)
		copy(dAtA[i:], m.EntityID)
		i = encodeVarintNetdemo(dAtA, i, uint64(len(m.EntityID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintNetdemo(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNetdemo(dAtA []byte, offset int, v uint64) int {
	offset -= sovNetdemo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UdpAddr)
	if l > 0 {
		n += 1 + l + sovNetdemo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegisterResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result {
		n += 2
	}
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovNetdemo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CreateEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovNetdemo(uint64(l))
	}
	l = len(m.EntityID)
	if l > 0 {
		n += 1 + l + sovNetdemo(uint64(l))
	}
	if m.EntityType != 0 {
		n += 1 + sovNetdemo(uint64(m.EntityType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DestroyEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovNetdemo(uint64(l))
	}
	l = len(m.EntityID)
	if l > 0 {
		n += 1 + l + sovNetdemo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PosAndYaw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovNetdemo(uint64(l))
	}
	l = len(m.EntityID)
	if l > 0 {
		n += 1 + l + sovNetdemo(uint64(l))
	}
	if m.PosX != 0 {
		n += 5
	}
	if m.PosY != 0 {
		n += 5
	}
	if m.PosZ != 0 {
		n += 5
	}
	if m.Yaw != 0 {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNetdemo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNetdemo(x uint64) (n int) {
	return sovNetdemo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterClient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetdemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UdpAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetdemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetdemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UdpAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetdemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetdemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Result = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetdemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetdemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetdemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetdemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetdemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetdemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetdemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetdemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityType", wireType)
			}
			m.EntityType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EntityType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNetdemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DestroyEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetdemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DestroyEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DestroyEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetdemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetdemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetdemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetdemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetdemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PosAndYaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetdemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PosAndYaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PosAndYaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetdemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetdemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetdemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetdemo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field PosX", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.PosX = float32(math.Float32frombits(v))
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field PosY", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.PosY = float32(math.Float32frombits(v))
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field PosZ", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.PosZ = float32(math.Float32frombits(v))
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Yaw", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Yaw = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipNetdemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetdemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNetdemo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetdemo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetdemo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNetdemo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNetdemo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNetdemo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNetdemo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetdemo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNetdemo = fmt.Errorf("proto: unexpected end of group")
)