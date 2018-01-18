package msgcenter

import (
	"net"
	"sync"
	"strconv"
	"time"
	"fmt"
	"peenet/memery"
	"peenet/task"
	"peenet/common"
	"peenet/config"
	"log"
	"peenet/push"
)

type IDial interface {
	Start()
	Close()
	Send(data string)
}

type InsideDial struct {

	IInsideProto

	wsWorkId  		int64
	conn 			*net.UDPConn
	syncWait 		sync.WaitGroup
	lockMemory		memery.IMemory
	insideTask		task.ITask
	deviceId 		string
	deviceStus		int
	deviceCstu		int
}

func NewInsideDial(workId int64, memory memery.IMemory, tasks task.ITask) *InsideDial{
	return &InsideDial{
		wsWorkId:workId,
		lockMemory:memory,
		insideTask:tasks,
	}
}

func (udp *InsideDial) Start(){

	udp.deviceId = common.GetRandNum(19)

	udp.lockMemory.Set("wt:ws:inside:" + strconv.FormatInt(udp.wsWorkId,10), udp.deviceId, time.Hour * 24 * 12 * 2)

	port, _ := strconv.Atoi(conf.PeeConf["server_udp_port"])

	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.ParseIP(conf.PeeConf["server_udp_addr"]),
		Port: port,
	})

	fmt.Println(net.ParseIP(conf.PeeConf["server_udp_addr"]), port)


	if err != nil {
		log.Println("websocket inside udp server connect fail - error from instanceUDP")
		fmt.Println(err)
		return
	}

	fmt.Println(udp.deviceId, err)

	udp.conn = socket

	go udp.pipe()

	udp.initDeviceData()

	go func() {
		ticker := time.NewTicker(time.Second * 30)

		for  {
			select {
			case <- ticker.C:
				udp.Send("wthblv" + strconv.Itoa(udp.deviceCstu) + udp.deviceId)
			}
		}
	}()


}

func (udp *InsideDial) initDeviceData(){
	udp.deviceStus, udp.deviceCstu = 80, 4
}

func (udp *InsideDial) Close(){

	udp.conn.Close()

	udp.insideTask.PushMessage(common.NewComposition_UDP(udp.wsWorkId, "close", common.Close))
}

func (udp *InsideDial) Send(data string){
	udp.conn.Write([]byte(data))
}

func (udp *InsideDial) MessageDeal(data string){
	switch data {
	case "wtopen1":
	case "wtopen3":
		udp.deviceCstu = 3
		break
	case "wtclse1":
	case "wtclse3":
		udp.deviceCstu = 4
		break
	default:
		log.Printf("deviceId: %s 收到未知消息！！", data)
		break
	}
}

func (udp *InsideDial) pipe(){

	for  {
		data := make([]byte, 1024)
		index, _, err := udp.conn.ReadFromUDP(data)

		if err != nil {
			udp.Close()

			log.Println("udp server connect break off", err)
			break
		}

		msg := string(data[0:index])

		udp.insideTask.PushMessage(push.BuildWsInsideData(udp.wsWorkId, msg, common.Notice, common.Ws_Proto_Udp))
	}
}

