package storage

import (
	"io"
	"crypto/rand"
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
	"sync"
	"errors"
)

type DeviceDetail struct {
	Cstu int
	Stus int
	IP string
	Port int
	Guid string
}

var deviceList map[string]*DeviceDetail

var devices []string


type deviceHandler struct {
	syncTex sync.RWMutex
}


func (device *deviceHandler) CreateDevice() *DeviceDetail {

	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	guid := GetMd5String(base64.URLEncoding.EncodeToString(b))


	detail := &DeviceDetail{
		Cstu: 3,
		Stus:80,
		Guid:guid,
	}

	deviceList[guid] = detail

	devices = append(devices, guid)

	return detail
}


func (device *deviceHandler) RemoveDevice(deviceId string){

	device.syncTex.Lock()

	defer device.syncTex.Unlock()

	delete(deviceList, deviceId)

	devices = RemoveSlice(devices, deviceId)

}

func (device *deviceHandler) GetDeviceDetail(deviceId string)(*DeviceDetail, error) {

	if v, ok := deviceList[deviceId]; ok {
		return v, nil
	}

	return nil, errors.New("not exits")
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func RemoveSlice(slice []string, elems ...string) []string {
	isInElems := make(map[string]bool)
	for _, elem := range elems {
		isInElems[elem] = true
	}
	w := 0
	for _, elem := range slice {
		if !isInElems[elem] {
			slice[w] = elem
			w += 1
		}
	}
	return slice[:w]
}


func NewDeviceDetail() *DeviceDetail{
	return &DeviceDetail{
	}
}

func NewDeviceHandler() *deviceHandler{

	return &deviceHandler{}
}


