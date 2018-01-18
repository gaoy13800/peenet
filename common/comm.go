package common

import (
	"time"
)

//协议类型
type ProtoT uint
const (
	STCP ProtoT = iota + 1
	SUDP
	SWS
	SHTTP

	CTCP
	CUDP
	CHTTP
)

//动作类型
type Action uint
const (
	Connect Action = iota + 1
	Close
	Notice

	Ping
	Pong
	Tester

)

// 会话类型
type SessionT uint
const (
	SWEB SessionT = iota + 1
	SSocket
	SWEBSOCKET

	CWEBSOCKET
	CSOCKET
)

//主程消息类型
type MessageT uint
const (
	Message_ws MessageT = iota + 1
	Message_http
	Message_tcp
	Message_udp
)

//webSocket 内部协议类型
type WsInsideT uint
const (
	Ws_Proto_Udp WsInsideT = iota + 1
	Ws_Proto_Tcp
	Ws_Proto_Http
)

const Ws_Query = "query"
const Ws_Talk = "talk"



// webSocket 内部udp队列传输结构
type WsUDP_Composition struct {
	WsWorkId	int64
	Data string
	MessageType Action
}

func NewComposition_UDP(workId int64, data string, action Action) *WsUDP_Composition{
	return &WsUDP_Composition{
		WsWorkId:workId,
		Data:data,
		MessageType:action,
	}
}


type GlobalMemoryT uint

const (
	Global_Http GlobalMemoryT = iota + 1
	Global_Ws
	Global_Socket
	Global_Other
	Global  // 真正的全局库无论从哪里拿都是一个库

	User_Library
	Common_Library
)

const LOGIN_COMMON_PASSWD = "111111"

const LongTimeMemory = time.Hour * 24 * 12 * 2

