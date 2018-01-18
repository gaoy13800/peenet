package msgcenter

import (
	"fmt"
	"net"
	"sync"
	"log"
	"peenet/task"
	"peenet/common"
	"peenet/config"
	"peenet/push"
	"peenet/memery"
)


//新版tcp协议


/**
	wtgoid + deviceId   Connect
	wtopen				Open
	wtclse				Close
	wtbrut				Brut
 */



type IInsideDialTCP interface {
	Start()

	Close()

	RawSend(data string)
}

type InsideTcp struct {

	InsideProto

	wsWorkId int64

	conn 	 *net.TCPConn

	syncWait	sync.WaitGroup

	deviceId 	string

	insideTask task.ITask

	memory      memery.IMemory

	deviceStatus string
	deviceElec	 string
}

func (tcp *InsideTcp) Start(){

	tcp.initDevice()

	tcp.deviceId = common.GetRandNum(15)

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", conf.PeeConf["server_dial_port"]))

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTCP("tcp", nil,  tcpAddr)

	if err != nil {
		fmt.Println(err)
		return
	}

	tcp.conn = conn

	go tcp.pipe()

}

func (tcp *InsideTcp) Close(){
	tcp.conn.Close()
}

func (tcp *InsideTcp) RawSend(data string){

	tcp.conn.Write([]byte(data))
}

func (tcp *InsideTcp) pipe(){
	for  {
		byt := make([]byte, 1024)

		index, err := tcp.conn.Read(byt)

		if err != nil {

			log.Println("webSocket 内部tcp连接中断！！！" , err)

			tcp.Close()

			break
		}


		msg := string(byt[0:index])

		log.Println("接收到tcp服务器端信息", msg)

		tcp.insideTask.PushMessage(push.BuildWsInsideData(tcp.wsWorkId, msg, common.Notice, common.Ws_Proto_Tcp))
	}
}

func (tcp *InsideTcp) initDevice(){
	tcp.deviceStatus, tcp.deviceElec = "4", "80"
}

func NewInsideTcp(workId int64, task task.ITask, memory memery.IMemory)*InsideTcp{

	return &InsideTcp{
		wsWorkId:workId,
		insideTask:task,
		memory:memory,
	}

}


