package storage


type UserStorage struct {
	UserIds []int
	DeviceIds []string
	DeviceMaps map[string]*DeviceDetail
}

