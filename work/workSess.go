package work

import (
	"sync"
	"peenet/common"
)

/**
	session centralized processing
	session 集中处理
 */

type ISession interface {
	ID() int64

	DeviceId() string
}

type SessionComm struct {

	deviceId string //设备id

	WithClose func()

	id int64

	wait sync.WaitGroup

}

func (sess *SessionComm) ID() int64{
	return sess.id
}

func (sess *SessionComm) DeviceId() string{
	return sess.deviceId
}

func (sess *SessionComm) SetDeviceId(id string){
	sess.deviceId = id
}

func NewSession(t common.SessionT)interface{}{
	switch t {
	case common.SWEB:
		return new(Session_Ws)
		break
	case common.SSocket:
		return new(Session_Ws)
		break
	case common.SWEBSOCKET:
		return new(Session_Ws)
		break
	default:
		break
	}

	return nil
}



