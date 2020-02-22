/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.16
 * Note:
 */

package connectorservice

import (
	"github.com/longhf/gosim/common"
	"github.com/longhf/gosim/network/config"
	"github.com/longhf/gosim/network/proxy"
	hm "github.com/longhf/gosim/network/netutil/handlermanager"
	"net"
)

type ConnectorService struct {
	clientId common.ClientId
	dialAddr string
	udpSrcAddr string
	serverProxy *proxy.ServerProxy
	udpServerProxy *proxy.ConnectedUdpProxy
	handlerManager *hm.HandlerManager

	onCreateEntity common.CreateEntityHandleFunc
	onDestroyEntity common.DestroyEntityHandleFunc
	onPosAndYaw common.PosAndYawHandleFunc
}


func NewConnectorService() *ConnectorService {
	return &ConnectorService{
		handlerManager: hm.NewHandlerManager(),
	}
}


func (cs *ConnectorService) Run(dialAddr string, udpSrcAddr string) {
	cs.dialAddr = dialAddr
	cs.udpSrcAddr = udpSrcAddr
	cs.registerHandler()
	go cs.dialUDP()
	cs.dialTCP()
}


func (cs *ConnectorService) dialTCP() error {
	conn, err := net.Dial("tcp", cs.dialAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetReadBuffer(config.SERVER_PROXY_READ_BUFFER_SIZE)
	tcpConn.SetWriteBuffer(config.SERVER_PROXY_WRITE_BUFFER_SIZE)
	tcpConn.SetNoDelay(config.SERVER_PROXY_SET_NO_DELAY)
	cs.serverProxy = proxy.NewServerProxy(conn)
	return cs.serverProxy.Run(cs.handlerManager)
}


func (cs *ConnectorService) dialUDP() error {
	srcAddr, err := net.ResolveUDPAddr("udp", cs.udpSrcAddr)
	if err != nil {
		return err
	}
	dstAddr, err := net.ResolveUDPAddr("udp", cs.dialAddr)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err!= nil {
		return err
	}
	defer conn.Close()
	conn.SetReadBuffer(config.UDP_SERVER_PROXY_READ_BUFFER_SIZE)
	conn.SetWriteBuffer(config.UDP_SERVER_PROXY_WRITE_BUFFER_SIZE)
	cs.udpServerProxy = proxy.NewConnectedUdpProxy(*conn)
	return cs.udpServerProxy.Run(cs.handlerManager)
}


// #-------------------- Handle Msg ----------------------------
func (cs *ConnectorService) OnCreateEntity(f common.CreateEntityHandleFunc) {
	cs.onCreateEntity = f
}

func (cs *ConnectorService) OnDestroyEntity(f common.DestroyEntityHandleFunc) {
	cs.onDestroyEntity = f
}

func (cs *ConnectorService) OnPosAndYaw(f common.PosAndYawHandleFunc) {
	cs.onPosAndYaw = f
}

// ------------------*- Handler -*--------------------
func (cs *ConnectorService) registerHandler() {
	register := cs.msgHandlerManager.RegisterMsgHandler
	register(proto.MT_REGISTER_RESULT, &proto.RegisterResult{}, cs.HandleRegisterResult)
	register(proto.MT_CREATE_ENTITY, &proto.CreateEntity{}, cs.HandleCreateEntity)
	register(proto.MT_DESTROY_ENTITY, &proto.DestroyEntity{}, cs.HandleDestroyEntity)
	register(proto.MT_POS_AND_YAW, &proto.PosAndYaw{}, cs.HandlePosAndYaw)
}


func (cs *ConnectorService) HandleRegisterResult(msg interface{}) {
	rr := msg.(*proto.RegisterResult)
	result := rr.Result
	if !result {
		// TODO
		fmt.Println("Client Register Fail..")
		return
	}
}


func (cs *ConnectorService) HandleCreateEntity(msg interface{}) {
	ce := msg.(*proto.CreateEntity)
	entityId := common.EntityId(ce.EntityID)
	entityType := common.EntityType(ce.EntityType)
	if cs.onCreateEntity != nil {
		cs.onCreateEntity(entityId, entityType)
	}
}

func (cs *ConnectorService) HandleDestroyEntity(msg interface{}) {
	de := msg.(*proto.DestroyEntity)
	entityId := common.EntityId(de.EntityID)
	if cs.onDestroyEntity != nil {
		cs.onDestroyEntity(entityId)
	}
}

func (cs *ConnectorService) HandlePosAndYaw(msg interface{}) {
	pAy := msg.(*proto.PosAndYaw)
	entityId := common.EntityId(pAy.EntityID)
	if cs.onPosAndYaw != nil {
		cs.onPosAndYaw(entityId, pAy.PosX, pAy.PosY, pAy.PosZ, pAy.Yaw)
	}
}

// #-------------------- Send Msg ----------------------------
func (cs *ConnectorService) SendCreateEntity(entityId common.EntityId, entityType common.EntityType) {
	cs.serverProxy.SendCreateEntity(cs.clientId, entityId, entityType)
}


func (cs *ConnectorService) SendDestroyEntity(entityId common.EntityId) {
	cs.serverProxy.SendDestroyEntity(cs.clientId, entityId)
}


func (cs *ConnectorService) SendPosAndYaw(entityId common.EntityId, x, y, z, yaw float32) {
	cs.udpServerProxy.SendPosAndYaw(cs.clientId, entityId, x, y, z, yaw)
}
