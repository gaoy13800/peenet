package msgcenter

import (
	"sync"
	"fmt"
)

type PeeSocketMessage struct {

	syncTex 	sync.RWMutex

	deviceHeart string
}



func (this *PeeSocketMessage)Dispense(msg string){


	fmt.Println("socket tcp recevi message :", msg)

}


func NewSocketMessage()*PeeSocketMessage{

	return &PeeSocketMessage{

	}

}