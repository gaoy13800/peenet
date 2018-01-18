package push

import "peenet/common"

type AnyData struct {
	T common.MessageT
	Sess interface{}
	Data string
	Action common.Action
	Misc interface{}
	InsideType string
}

func Build(t common.MessageT, action common.Action, sess interface{}, data string, misc interface{}, insideType string) *AnyData{
	return &AnyData{
		T: t,
		Sess:sess,
		Data:data,
		Action:action,
		Misc:misc,
		InsideType:insideType,
	}
}



type WsInsideData struct {

	WorkId int64

	Data    string

	Action   common.Action

	InsideType common.WsInsideT
}

func BuildWsInsideData(workId int64, data string, action common.Action, t common.WsInsideT)*WsInsideData{

	return &WsInsideData{
		WorkId: workId,
		Data: data,
		Action:action,
		InsideType:t,
	}

}