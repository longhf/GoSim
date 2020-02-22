/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.16
 * Note:
 */

package gateservice

import (
	"github.com/longhf/gosim/common"
	"github.com/longhf/gosim/network/config"
	"github.com/longhf/gosim/network/proxy"
	hm "github.com/longhf/gosim/network/netutil/handlermanager"
	"net"
)


type GateService struct {
	listenAddr string
	udpProxy *proxy.UdpProxy
	clientProxyMap map[common.ClientId]*proxy.ClientProxy
	handlerManager *hm.HandlerManager

	onCreateEntity common.CreateEntityHandleFunc
	onDestroyEntity common.DestroyEntityHandleFunc
	onPosAndYaw common.PosAndYawHandleFunc
}


func NewGateService() *GateService {
	return &GateService{
		clientProxyMap: make(map[common.ClientId]*proxy.ClientProxy, 0),
		handlerManager: hm.NewHandlerManager()
	}
}


func (gs *GateService) Run(listenAddr string) {
	gs.listenAddr = listenAddr
	gs.registerHandler()
	go common.RunForever(config.RESTART_UDP_proxy_INTERVAL, gs.serveUDP)
	common.RunForever(config.RESTART_TCP_proxy_INTERVAL, gs.serveTcp)
}


func (gs *GateService) serveTcp() error {
	lis, err := net.Listen("tcp", gs.listenAddr)
	if err != nil {
		return err
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil{
			return err
		}
		go gs.handleTcpConnection(conn)
	}
}


func (gs *GateService) handleTcpConnection(conn net.Conn){
	clientId := common.GenClientId()
	tcpConn := conn.(*net.TcpConn)
	tcpConn.SetReadBuffer(config.CLIENT_PROXY_READ_BUFFER_SIZE)
	tcpConn.SetWriteBuffer(config.CLIENT_PROXY_WRITE_BUFFER_SIZE)
	tcpConn.SetNoDelay(config.CLIENT_PROXY_SET_NO_DELAY)
	clientProxy := proxy.NewClientProxy(conn, clientId)
	gs.clientProxyMap[clientId] = clientProxy
	clientProxy.Run(gs.handlerManager)
}


func (gs *GateService) serveUDP() error {
	listenAddr, err := net.ResolveUDPAddr("udp", gs.listenAddr)
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		return err
	}
	defer udpConn.Close()
	udpConn.SetReadBuffer(config.UDP_CLIENT_PROXY_READ_BUFFER_SIZE)
	udpConn.SetWriteBuffer(config.UDP_CLIENT_PROXY_WRITE_BUFFER_SIZE)
	gs.udpProxy = proxy.NewUdpProxy(*udpConn)
	return gs.udpProxy.Run(gs.handlerManager)
}

// #-------------------- Handle Msg ----------------------------
func (gs *GateService) OnCreateEntity(f common.CreateEntityHandleFunc) {
	gs.onCreateEntity = f
}


func (gs *GateService) OnDestroyEntity(f common.DestroyEntityHandleFunc) {
	gs.onDestroyEntity = f
}


func (gs *GateService) OnPosAndYaw(f common.PosAndYawHandleFunc) {
	gs.onPosAndYaw = f
}

// ------------------*- Handler -*--------------------
func (gs *GateService) registerHandler() {
	register := gs.handlerManager.RegisterHandler
	register(proto.MT_REGISTER_CLIENT, &proto.RegisterClient{}, gs.HandleRegisterClient)
	register(proto.MT_CREATE_ENTITY, &proto.CreateEntity{}, gs.HandleCreateEntity)
	register(proto.MT_DESTROY_ENTITY, &proto.DestroyEntity{}, gs.HandleDestroyEntity)
	register(proto.MT_POS_AND_YAW, &proto.PosAndYaw{}, gs.HandlePosAndYaw)
}


func (gs *GateService) HandleRegisterClient(msg interface{}) {
	fmt.Println("HandleRegisterClient")
	rc := msg.(*proto.RegisterClient)
	addr, err := net.ResolveUDPAddr("udp", rc.UdpAddr)
	if err != nil {
		return
	}
	gs.UdpAddr = addr
	gs.SendRegisterResult(gs.clientId, true)
}


func (gs *GateService) HandleCreateEntity(msg interface{}) {
	var entityId common.EntityId
	ce := msg.(*proto.CreateEntity)
	if ce.EntityID != "" {
		entityId = common.EntityId(ce.EntityID)
	} else {
		entityId = common.GenEntityId()
	}
	entityType := common.EntityType(ce.EntityType)
	if gs.onCreateEntity != nil {
		gs.onCreateEntity(entityId, entityType)
	}
}


func (gs *GateService) HandleDestroyEntity(msg interface{}) {
	de := msg.(*proto.DestroyEntity)
	entityId := common.EntityId(de.EntityID)
	if gs.onDestroyEntity != nil {
		gs.onDestroyEntity(entityId)
	}
}


func (gs *GateService) HandlePosAndYaw(msg interface{}) {
	pAy := msg.(*proto.PosAndYaw)
	entityId := common.EntityId(pAy.EntityID)
	if gs.onPosAndYaw != nil {
		gs.onPosAndYaw(entityId, pAy.PosX, pAy.PosY, pAy.PosZ, pAy.Yaw)
	}
}

// #-------------------- Send Msg ----------------------------
func (gs *GateService) SendCreateEntity(clientIds []common.ClientId, entityId common.EntityId, entityType common.EntityType) {
	length := len(clientIds)
	if length < 1 {
		// Broadcast
		for cid, gs := range gs.clientProxyMap {
			gs.SendCreateEntity(cid, entityId, entityType)
		}
	} else {
		// Multicast
		for _, cid := range clientIds {
			if gs, ok := gs.clientProxyMap[cid]; ok {
				gs.SendCreateEntity(cid, entityId, entityType)
			}
		}
	}
}


func (gs *GateService) SendDestroyEntity(clientIds []common.ClientId, entityId common.EntityId) {
	length := len(clientIds)
	if length < 1 {
		// Broadcast
		for cid, gs := range gs.clientProxyMap {
			gs.SendDestroyEntity(cid, entityId)
		}
	} else {
		// Multicast
		for _, cid := range clientIds {
			if gs, ok := gs.clientProxyMap[cid]; ok {
				gs.SendDestroyEntity(cid, entityId)
			}
		}
	}
}


func (gs *GateService) SendPosAndYaw(clientIds []common.ClientId, entityId common.EntityId, x, y, z, yaw float32) {
	length := len(clientIds)
	if length < 1 {
		// Broadcast
		for cid, gs := range gs.clientProxyMap {
			// gs.SendPosAndYaw(cid, entityId, x, y, z, yaw)
			gs.udpProxy.SendPosAndYaw(gs.UdpAddr, cid, entityId, x, y, z, yaw)
		}
	} else {
		// Multicast
		for _, cid := range clientIds {
			if gs, ok := gs.clientProxyMap[cid]; ok {
				// gs.SendPosAndYaw(cid, entityId, x, y, z, yaw)
				gs.udpProxy.SendPosAndYaw(gs.UdpAddr, cid, entityId, x, y, z, yaw)
			}
		}
	}
}
