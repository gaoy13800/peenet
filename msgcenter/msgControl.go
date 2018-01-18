package msgcenter

import (
	"log"
	"peenet/common"
	"fmt"
	"peenet/push"
	"peenet/work"
)

type ICenter interface {

}

type CenterMessage struct {
	wsCenter	*PeeWsMsg
}

func (msg *CenterMessage) Distribute(pushData *push.AnyData){

	t, data, session, workId,insideType := pushData.T, pushData.Data, pushData.Sess.(work.ISession), pushData.Misc, pushData.InsideType

	log.Println("消息中心收到message:", data)

	if data == "" {
		return
	}

	switch t {
	case common.Message_http:
		fmt.Println(data)
		break
	case common.Message_ws:
		msg.wsCenter.Dispense(data, session, workId.(int64), insideType)
		break
	case common.Message_tcp:
		fmt.Println("tcp")
		break
	default:
		fmt.Println("unknown message type, rejected processing")
		return
	}

}

func (msg *CenterMessage) Focus(){

}

func NewCenterMessage() *CenterMessage{

	return &CenterMessage{
		wsCenter:NewWsMsg(),
	}

}