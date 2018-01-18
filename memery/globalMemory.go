package memery

import (
	"time"
	"peenet/common"
)

/**
	内存存储器
 */

var Http_Global = New(10 * time.Minute, 10 * time.Second)

var Ws_Global   = New(10 * time.Minute, 10 * time.Second)

var Any_Global  = New(10 * time.Minute, 10 * time.Second)

var New_Global  = New(10 * time.Minute, 10 * time.Second)

var Sock_Global = New(10 * time.Minute, 10 * time.Second)

var UserLibrary  = New(10 * time.Minute, 10 * time.Second)

var ShareLibrary = New(10 * time.Minute, 10 * time.Second)

func SelectMemory(t common.GlobalMemoryT) *Cache{
	switch t {
	case common.Global_Http:
		return Http_Global
	case common.Global_Ws:
		return Ws_Global
	case common.Global_Socket:
		return Sock_Global
	case common.Global_Other:
		return New_Global
	case common.User_Library:
		return UserLibrary
	case common.Common_Library:
		return ShareLibrary
	default:
		return Any_Global
	}
}

