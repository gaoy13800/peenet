package common

import (
	"time"
	"math/rand"
	"strconv"
)

func GetRandNum(index int) string{

	var str string

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i:=0; i < index; i++ {

		num := strconv.Itoa(r.Intn(10))

		str += num
	}

	return str
}