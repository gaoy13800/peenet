package work

/**
	webSocket session work interactive
 */

import (
	"github.com/gorilla/websocket"
	"sync"
	"log"
	"peenet/common"
	"net"
	"peenet/task"
	"peenet/push"
	"peenet/memery"
	"strconv"
	"time"
	"encoding/json"
)

type Session_Ws struct {

	SessionComm

	WithClose func()

	level	string // ws区分 会话级 查询级

	conn *websocket.Conn

	syncWait sync.WaitGroup

	udpClient net.UDPConn

	tasks task.ITask

	inside_type string
}

func (this *Session_Ws) Walk (conn interface{}, taskDis task.ITask){

	this.conn = conn.(*websocket.Conn)

	this.tasks = taskDis

	this.syncWait.Add(1)

	go func() {
		this.syncWait.Wait()

		this.Close()
	}()

	this.Encode(websocket.TextMessage, "hello peenet")

	go this.pipeReceive()
}

/**
	ws receive 接受消息
 */
func (this *Session_Ws) pipeReceive(){

	for {
		data, err := this.Decode()

		log.Println(data)

		if err != nil {
			log.Println("read:", err)
			goto WSCLOSE
		}

		if this.level == common.Ws_Query {
			this.queryHandler()

			log.Println("websocket 查询级收到:", data)
			continue
		}

		this.tasks.PushMessage(push.Build(common.Message_ws, common.Notice, this, data, this.id, this.inside_type)) //消息通知
		continue

	WSCLOSE:
		this.tasks.PushMessage(push.Build(common.Message_ws, common.Close, this, data, this.id, this.inside_type)) // 关闭通知
		break
	}
}


/**
	查询级ws
 */
func (this *Session_Ws) queryHandler(){

	commonMemory := memery.SelectMemory(common.Common_Library)

	resultMap := make(map[string]interface{})

	if count, ok := commonMemory.Get("user:id"); ok{
		resultMap["onlineNumber"] = count
	}else {
		resultMap["onlineNumber"] = 0
	}

	result, _ := json.Marshal(resultMap)

	this.Encode(websocket.TextMessage, string(result))
}

/**
	ws send 发送消息
 */
func (this *Session_Ws) Encode(mt int, message string ) bool{

	err := this.conn.WriteMessage(mt, []byte(message))

	if err != nil {
		log.Println("write:", err)
		return false
	}

	return true
}

/**
	读取消息
 */
func (this *Session_Ws) Decode() (string, error){
	mt, message, err := this.conn.ReadMessage()

	if err != nil || mt != 1 {
		return "", err
	}

	return string(message), nil
}

/**
	会话关闭
 */
func (this *Session_Ws) Close(){
	this.conn.Close()
	this.WithClose()
}

func (this *Session_Ws) SetupRelevant(level string, userId, proto_type, deviceId string){

	this.level = level

	if level == common.Ws_Talk {

		this.inside_type = proto_type

		this.deviceId = deviceId

		memoryCommon := memery.SelectMemory(common.Common_Library)

		if _, ok := memoryCommon.Get("user:count"); !ok {
			memoryCommon.Set("user:count", 0, common.LongTimeMemory)
		}

		memoryCommon.Increment("user:count", 1)

		memoryCommon.Set("ws:work:device:" + strconv.FormatInt(this.id,10), deviceId, common.LongTimeMemory)
	} else {

		go func() {
			tick := time.Tick(time.Second * 50)

			for  {
				select {
				case <- tick:
					this.queryHandler()
				}
			}
		}()

	}
}

/**
    获取回话内部协议类型
 */
func (this *Session_Ws) ProtoType() string{
	return this.inside_type
}

/**
	获取会话协议类型
 */
func (this *Session_Ws) SessionType() common.SessionT{
	return common.SWEBSOCKET
}