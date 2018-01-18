package msgcenter

import (
	"sync"
	"log"
	"peenet/memery"
	"peenet/common"
	"peenet/work"
	"peenet/task"
	"peenet/push"
	"github.com/gorilla/websocket"
)

type PeeWsMsg struct {
	syncTex 	sync.RWMutex

	peeMemory   memery.IMemory

	tasks  		task.ITask

	insideMng	IInsideManage
}

func (wsMsg *PeeWsMsg) Dispense(msg string, session work.ISession, workId int64, clientType string){

	inside := wsMsg.insideMng.Get(workId)

	if inside == nil {
		switch clientType {
		case "tcp":
			go wsMsg.InstanceTCP(msg, workId, wsMsg.peeMemory, wsMsg.tasks, session)
			return
		case "udp":
			go wsMsg.InstanceUDP(msg, workId, wsMsg.peeMemory, wsMsg.tasks, session)
			return
		default:
			break
		}

		return
	}

	switch inside.(type) {

	case IInsideDialTCP:
		inside.(*InsideTcp).RawSend(msg)
		return
	case IDial:
		inside.(IDial).Send(msg)
		return
	default:
		log.Println("unkown inside type")

		return
	}
}

/**
	初始化tcp连接
 */
func (wsMsg *PeeWsMsg) InstanceTCP(data string, workId int64, memory memery.IMemory, iTask task.ITask, sess work.ISession){

	dial := NewInsideTcp(workId, iTask, memory)

	dial.Start()

	wsMsg.insideMng.Set(workId, dial)

	wsMsg.ReceiveTask(sess, iTask)
}

func (wsMsg *PeeWsMsg) InstanceUDP(data string, workId int64, memory memery.IMemory, tasks task.ITask, sess work.ISession){

	dialSession :=  NewInsideDial(workId, memory, tasks)

	dialSession.Start()

	wsMsg.insideMng.Set(workId, dialSession)

	wsMsg.ReceiveTask(sess, tasks)
}

func (wsMsg *PeeWsMsg) ReceiveTask(sess work.ISession, ITask task.ITask){


	go func() {

		for  {
			msg := <- ITask.GetTask()

			_, action, insideMessage := msg.(*push.WsInsideData).InsideType, msg.(*push.WsInsideData).Action, msg.(*push.WsInsideData).Data

			if action == common.Notice{
				sess.(*work.Session_Ws).Encode(websocket.TextMessage, insideMessage)
			}else if action == common.Close{
				sess.(*work.Session_Ws).Encode(websocket.TextMessage, "内部会话结束 请重新连接")
			}
		}
	}()

	/*for  {


		switch meType {
		case common.Ws_Proto_Tcp:
			if action == common.Notice{
				fmt.Println("Encode Ws Message,", insideMessage)
				sess.(*work.Session_Ws).Encode(websocket.TextMessage, insideMessage)
			}else if action == common.Close{
				sess.(*work.Session_Ws).Encode(websocket.TextMessage, "内部会话结束 请重新连接")
			}
			return
		case common.Ws_Proto_Udp:
			if action == common.Notice{
				sess.(*work.Session_Ws).Encode(websocket.TextMessage, insideMessage)
			}else if action == common.Close {
				sess.(*work.Session_Ws).Encode(websocket.TextMessage, "内部会话结束 请重新连接")
			}
		}
	}*/
}


func NewWsMsg() * PeeWsMsg{

	insideTask := task.TaskCreate()

	return &PeeWsMsg{
		peeMemory:memery.NewMeInstance(),
		tasks:insideTask,
		insideMng:NewInsideManage(),
	}
}

