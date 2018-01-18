package work

import (
	"sync"
	"sync/atomic"
	"peenet/common"
)

const totalTryCount  = 10000

type SessionManager struct {

	sessionList  map[int64]ISession

	sessionIDAcc int64

	syncTex sync.RWMutex
}

func (this *SessionManager) ADD(sess ISession, t common.SessionT){

	this.syncTex.Lock()
	defer this.syncTex.Unlock()

	var tryCount int = totalTryCount

	var id int64

	for tryCount > 0 {
		id = atomic.AddInt64(&this.sessionIDAcc, 1)
		if _, ok := this.sessionList[id]; !ok {
			break
		}
		tryCount--
	}

	//socketSes := session.(*SocketSession)
	//socketSes.id = id

	if t == common.SWEBSOCKET{
		sess.(*Session_Ws).id = id
	}

	this.sessionList[id] = sess

}

func (this *SessionManager) Remove(session ISession){

	this.syncTex.Lock()
	defer this.syncTex.Unlock()

	delete(this.sessionList, session.ID())
}

func (this *SessionManager) SessionCount() int{

	this.syncTex.Lock()
	defer this.syncTex.Unlock()

	return len(this.sessionList)
}

func (this *SessionManager) RemoveAll(){

	for _,sess := range this.sessionList{

		this.Remove(sess)
	}

}


func NewSessionManage () *SessionManager{

	return &SessionManager{
		sessionList:make(map[int64]ISession),
	}

}

//func (sem *SessionManager) GetSingleSessionCount()(map[string]int){
//	var wsCount int
//
//	var res map[string]int
//
//	for _, sess := range sem.sessionList{
//
//		switch sess.(type) {
//
//		case Session_Ws:
//			wsCount++
//			break
//		default:
//			panic("session type invalid")
//			break
//		}
//	}
//
//	res["ws"] = wsCount
//
//	return res
//}

