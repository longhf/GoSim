/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.08
 * Note: Reflect 有性能代价，简化代码先写成动态的，若有性能瓶颈，再改为静态的
 */

package handlermanager

import (
	"errors"
	"fmt"
	pb "github.com/golang/protobuf/proto"
	"github.com/longhf/gosim/network/types"
	"reflect"
)

type msgHandler func(interface{})

type handlerInfo struct {
	mType    reflect.Type
	mHandler msgHandler
}

type HandlerManager struct {
	handlerMap map[types.MsgType]handlerInfo
}

func NewHandlerManager() *HandlerManager {
	return &HandlerManager{handlerMap: make(map[types.MsgType]handlerInfo, 0)}
}

func (hm *HandlerManager) RegisterHandler(msgType types.MsgType, msg interface{}, handler msgHandler) {
	info := handlerInfo{
		mType: reflect.TypeOf(msg),
		mHandler: handler,
	}
	hm.handlerMap[msgType] = info
}

func (hm *HandlerManager) HandleMsg(msgType types.MsgType, data []byte) error{
	if info, ok := hm.handlerMap[msgType]; ok {
		msg := reflect.New(info.mType.Elem()).Interface()
		err := pb.Unmarshal(data, msg.(pb.Message))
		if err != nil {
			return err
		}
		info.mHandler(msg)
		return err
	}
	return errors.New(fmt.Sprintf("Not Registed MsgType %d ", msgType))
}
