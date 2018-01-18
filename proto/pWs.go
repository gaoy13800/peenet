package proto

import (
	"flag"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"peenet/work"
	"peenet/common"
	"peenet/task"
	"peenet/push"

)

type WsProto struct {

	Protocol

	WsList map[string]*websocket.Conn

	taskDis task.ITask

}

var wsAddr = flag.String("wsAddress", "0.0.0.0:12000", "http service address")

var upGrader = websocket.Upgrader{}


/**
	webSocket 启动方法 监听ws请求
 */
func (ws *WsProto) Run(task task.ITask){
	ws.taskDis = task

	ws.SessionManager = work.NewSessionManage()

	flag.Parse()

	log.SetFlags(0)

	http.HandleFunc("/web", ws.Work_ws)
	http.HandleFunc("/web1", ws.Work_ws)

	// create new session 访问一次就是一个新的会话

	log.Fatal(http.ListenAndServe(*wsAddr, nil))
}

/**
	请求ws回调的方法
 */
func (ws *WsProto) Work_ws(w http.ResponseWriter, r *http.Request){

	//important step 删除origin 因为gorilla有一个校验origin的方法是有问题的

	r.Header.Del("Origin")

	conn, err := upGrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	userId := r.FormValue("userId")

	if userId == "" {
		return
	}

	go ws.Handler(conn, r)
}

/**
	生成会话处理，抛送ws connect信息
 */
func (ws *WsProto) Handler(conn *websocket.Conn, request *http.Request){

	sess := work.NewSession(common.SWEBSOCKET).(*work.Session_Ws)

	sess.Walk(conn, ws.taskDis)

	ws.ADD(sess, common.SWEBSOCKET)

	sess.WithClose = func() {
		ws.Remove(sess)
	}


	pro_type := request.FormValue("proto_type")

	if request.FormValue("level") == common.Ws_Query {
		//查询维度的ws不需要通知消息中心处理消息

		sess.SetupRelevant(common.Ws_Query, request.FormValue("userId"), pro_type, "")
	}else {
		sess.SetupRelevant(common.Ws_Talk, request.FormValue("userId"), pro_type, request.FormValue("deviceId"))

		ws.taskDis.PushMessage(push.Build(common.Message_ws, common.Connect, sess, "connect", sess.ID(), pro_type))
	}


}

//广播
func (ws *WsProto) Broadcast(msg []byte){

	allConn := ws.WsList

	for _, single := range allConn{
		//single.
		single.WriteMessage(websocket.TextMessage, msg)
	}

}



