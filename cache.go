package cast

import (
	"sync"

	c "github.com/flywithbug/cache"
)

type CacheModel struct {
	Cache    *c.Cache
	Models   map[string]string
	Apis     map[string]string
	Lock     sync.RWMutex
	ErrorApi map[int]interface{}
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

func (ca *CacheModel) flush() {
	ca.Cache.Flush()
	ca.Models = make(map[string]string)
	ca.Apis = make(map[string]string)
	ca.ErrorApi = make(map[int]interface{})
}

func (ca *CacheModel) exist(key string) bool {
	_, b := ca.Cache.Get(key)
	return b
}

func (ca *CacheModel) setApi(key string) {
	ca.Lock.Lock()
	ca.Apis[key] = key
	ca.Lock.Unlock()
}

func (ca *CacheModel) setModel(key string) {
	ca.Lock.Lock()
	ca.Models[key] = key
	ca.Lock.Unlock()
}

func (ca *CacheModel) set(key string, value interface{}) {
	ca.Cache.Set(key, value, c.DefaultExpiration)
}

func (ca *CacheModel) get(key string) (interface{}, bool) {
	return ca.Cache.Get(key)
}

func (ca *CacheModel) delete(Key string) {
	ca.Cache.Delete(Key)
}

func (ca *CacheModel) allItems() map[string]c.Item {
	return ca.Cache.Items()
}

func (ca *CacheModel) setErrorMap(key int, value interface{}) {
	ca.ErrorApi[key] = value
}
