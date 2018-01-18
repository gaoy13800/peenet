package msgcenter

import (
	"sync"
	"peenet/common"
	"log"
)

type IInsideManage interface {

	Get(id int64) IInsideProto

	Set(id int64, session IInsideProto)

	Remove(id int64)

	Count() int
}

type insideManage struct {

	WsInsideList	map[int64]IInsideProto

	syncTex 		sync.RWMutex
}

func (inside *insideManage) Get(id int64) IInsideProto {

	inside.syncTex.Lock()
	defer inside.syncTex.Unlock()

	v, ok := inside.WsInsideList[id]

	if ok {
		return v
	}

	return nil
}

func (inside *insideManage) Set(id int64, session IInsideProto){

	inside.syncTex.Lock()
	defer inside.syncTex.Unlock()

	inside.WsInsideList[id] = session
}

func (inside *insideManage) Remove(id int64){
	inside.syncTex.Lock()
	defer inside.syncTex.Unlock()

	delete(inside.WsInsideList, id)
}

func (inside *insideManage) Count()int{
	return len(inside.WsInsideList)
}

func NewInsideManage() *insideManage{
	return &insideManage{
		WsInsideList:make(map[int64]IInsideProto),
	}
}



// ws内部公用类

type IInsideProto interface {
	Start()

	Close()

}

type InsideProto struct {

	wsWorkId int64

	deviceId string

}

func NewInsideProto(inside_type common.WsInsideT) IInsideProto{

	switch inside_type {
	case common.Ws_Proto_Tcp :
		return new(InsideTcp)
	case common.Ws_Proto_Udp:
		return new(InsideDial)
	default:
		log.Println("error ws inside type")
		return nil
	}

}