package memery

import (
	"time"
	"sync"
)


type IMemory interface {
	Set(key string, value interface{}, cacheTime time.Duration)

	Get(key string)(interface{}, bool)

	Count() int

	Delete(key string)
}

type meInstance struct {

	cacheSource *Cache

	syncTex sync.RWMutex
}


func (this *meInstance) Set(key string, value interface{}, cacheTime time.Duration){

	this.syncTex.Lock()

	defer this.syncTex.Unlock()

	this.cacheSource.Set(key, value, cacheTime)
}

func (this *meInstance) Get(key string)(interface{}, bool){
	this.syncTex.Lock()

	defer this.syncTex.Unlock()

	return this.cacheSource.Get(key)
}

func (this *meInstance) Count() int {
	return this.cacheSource.ItemCount()
}

func (this *meInstance) Delete(key string){
	this.syncTex.Lock()

	defer this.syncTex.Unlock()

	this.cacheSource.Delete(key)
}

func (this *meInstance) FlushAll(){
	this.cacheSource.Flush()
}

func NewMeInstance() *meInstance{

	instance := &meInstance{}

	instance.cacheSource = New(10 * time.Minute, 10 * time.Second)

	return instance
}
