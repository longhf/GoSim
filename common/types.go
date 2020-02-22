/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.06
 * Note:
 */

package common

import (
	"github.com/longhf/gosim/uuid"
)

// EntityID
type EntityId string

func (id EntityId) ToString() string{
	return string(id)
}

func (id EntityId) IsNil() bool{
	return id == ""
}

func GenEntityId() EntityId {
	return EntityId(uuid.GenUUID())
}


// ClientID
type ClientId string

func (id ClientId) ToString() string{
	return string(id)
}

func (id ClientId) IsNil() bool{
	return id == ""
}

func GenClientId() ClientId{
	return ClientId(uuid.GenUUID())
}


// EntityType
type EntityType uint32

const (
	ET_INVALID EntityType = iota
	ET_CAR
	ET_HUMAN
	ET_TRAFFIC_LIGHT
)


// HandleFunc
type CreateEntityHandleFunc func(EntityId, EntityType)
type DestroyEntityHandleFunc func(EntityId)
type PosAndYawHandleFunc func(EntityId, float32, float32, float32, float32)
