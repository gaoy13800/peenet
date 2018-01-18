package proto

import (
	"log"
	"net/http"
	peeHttp "peenet/http"
)

type HttpProto struct {
	Protocol
}

func (self *HttpProto) Run(){

	router := peeHttp.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func (self *HttpProto) Close(){


}