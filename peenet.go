package main

import (
	"peenet/common"
	"peenet/proto"
	"peenet/task"
	"peenet/msgcenter"
	"log"
	"peenet/push"
	"runtime"
	"fmt"
	"peenet/config"
)


var (

	wsTask task.ITask = task.TaskCreate()

	tcpTask task.ITask = task.TaskCreate()
)


func main(){


	banner := `

	 ____
    /    )
---/____/----__----__----__----__--_/_-
  /        /___) /___) /   ) /___) /
_/________(___ _(___ _/___/_(___ _(_ __

	`

	fmt.Println(banner)

	fmt.Println("[verison]", conf.PeeConf["version"])


	runtime.GOMAXPROCS(runtime.NumCPU())

	go proto.NewServer(common.SHTTP).(*proto.HttpProto).Run()

	go proto.NewServer(common.SWS).(*proto.WsProto).Run(wsTask)

	disposeMessage()

	log.Println("程序启动！！！")

	defer log.Println("程序结束！！！")

	select {}
}

func disposeMessage(){

	go func() {

		messageCenter := msgcenter.NewCenterMessage()

		for {
			taskData := <- wsTask.GetTask()

			messageCenter.Distribute(taskData.(*push.AnyData))
		}
	}()

}
