package cast

import (
	"sync"

	c "github.com/flywithbug/cache"
)

type CacheModel struct {
	Cache  *c.Cache
	Models map[string]string
	Apis   map[string]string
	Lock   sync.RWMutex
}

var ca *CacheModel

func init() {
	Cache().Cache = c.New(c.DefaultExpiration, c.DefaultExpiration)
}

func Cache() *CacheModel {
	if ca == nil {
		ca = new(CacheModel)
		ca.Cache = c.New(c.DefaultExpiration, c.DefaultExpiration)
		ca.Models = make(map[string]string)
		ca.Apis = make(map[string]string)
		ca.Lock = sync.RWMutex{}
	}
	return ca
}

func (ca *CacheModel) Flush() {
	ca.Cache.Flush()
	ca.Models = make(map[string]string)
	ca.Apis = make(map[string]string)
}

func (ca *CacheModel) Exist(key string) bool {
	_, b := ca.Cache.Get(key)
	return b
}

func (ca *CacheModel) SetApi(key string) {
	ca.Lock.Lock()
	ca.Apis[key] = key
	ca.Lock.Unlock()
}

func (ca *CacheModel) SetModel(key string) {
	ca.Lock.Lock()
	ca.Models[key] = key
	ca.Lock.Unlock()
}

func (ca *CacheModel) Set(key string, value interface{}) {
	ca.Cache.Set(key, value, c.DefaultExpiration)
}

func (ca *CacheModel) Get(key string) (interface{}, bool) {
	return ca.Cache.Get(key)
}

func (ca *CacheModel) Delete(Key string) {
	ca.Cache.Delete(Key)
}

func (ca *CacheModel) AllItems() map[string]c.Item {
	return ca.Cache.Items()
}
