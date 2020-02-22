/*
 * Copyright (c) 2019. Xiaolong Xu, ZhejiangLab.
 * Email: xuxiaolong@zhejianglab.com
 * Date: 2020.01.07
 * Note:
 */

package types

type MsgType uint16


const (
	MT_INVALID MsgType = iota
	MT_REGISTER_CLIENT
	MT_REGISTER_RESULT
	MT_CREATE_ENTITY
	MT_DESTROY_ENTITY
	MT_POS_AND_YAW
)
