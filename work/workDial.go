package work

import "net"

type DialSession struct {
	*SessionComm

	conn net.Conn

	deviceId string
}


func (cli *DialSession) Start(){



}

func (cli *DialSession) Stop(){

}

func (cli *DialSession) Walk(){

}

func (cli *DialSession) RawSend(data string){

}


