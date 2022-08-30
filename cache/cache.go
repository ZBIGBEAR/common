package cache

import (
	"sync"

	"github.com/BurntSushi/locker"
)

type ValueFunc func(string) (interface{}, error)

type cacheEntity struct {
	nlocker *locker.Locker
	mu      sync.Mutex
	cache   map[string]*entity
}

type entity struct {
	value interface{}
	err   error
}

var localCache *cacheEntity

func init() {
	localCache = &cacheEntity{
		cache: make(map[string]*entity),
	}
}

func Get(key string, f ValueFunc) (interface{}, error) {
	localCache.mu.Lock()
	defer localCache.mu.Unlock()

	v, ok := localCache.cache[key]
	if ok {
		return v.value, v.err
	}

	// 获取数据
	value, err := f(key)
	localCache.cache[key] = &entity{
		value: value,
		err:   err,
	}

	return value, err
}
