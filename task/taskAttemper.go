package task

import (
	"peenet/config"
	"strconv"
)

//任务调度器 buffer chan

type ITask interface {
	GetTask() chan interface{}
	PushMessage(anyData interface{})
}

type peeTask struct {
	Allot	chan interface{}
}

func (push *peeTask) GetTask() chan interface{} {
	return push.Allot
}

func (push *peeTask) PushMessage(anyData interface{}){
	push.Allot <- anyData
}

func TaskCreate() *peeTask {

	chanNum, _ := strconv.Atoi(conf.PeeConf["num"])

	this := &peeTask{
		Allot:make(chan interface{}, chanNum),
	}

	return this
}