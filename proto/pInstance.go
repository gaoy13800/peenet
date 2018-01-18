package proto

import (
	"net"
	"sync"
	"peenet/common"
	"errors"
	"peenet/work"
)

var (
	Unknown_Server_proto = errors.New("invalid server proto ")
)


type IProto interface {

	Close()

	Ping()

	Work()
}

type Protocol struct {
	conn net.Conn

	mutex sync.RWMutex

	wg sync.WaitGroup

	*work.SessionManager
}


func (proto *Protocol) Ping(){

}

func (proto *Protocol) Work(){

}

func (proto *Protocol) Close(){

}


func NewServer(t common.ProtoT)IProto{

	switch t {
	case common.SHTTP:
		return new(HttpProto)
	case common.SWS:
		return new(WsProto)
	default:
		panic(Unknown_Server_proto)
		break
	}

	return nil
}





