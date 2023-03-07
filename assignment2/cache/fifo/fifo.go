package fifo

import (
	"container/list"
	"lru/cache"
	"lru/cache/redis"
)

type FIFOCache struct {
	m              map[string]*list.Element
	cap            int
	l              list.List
	nextCacheLayer cache.Cache
}

func NewFIFOCache(cap int) cache.Cache {
	return &FIFOCache{
		m:              map[string]*list.Element{},
		cap:            cap,
		l:              list.List{},
		nextCacheLayer: redis.NewRedis(),
	}
}

type data struct {
	key   string
	value any
}

func (f *FIFOCache) Get(key string) any {
	el, ok := f.m[key]
	if !ok {
		val := f.nextCacheLayer.Get(key)
		f.Set(key, val)
		return val
	}
	return el.Value.(data).value
}

func (f *FIFOCache) Set(key string, value any) bool {
	if f.cap <= 0 {
		return false
	}
	el, ok := f.m[key]
	d := data{key: key, value: value}
	if !ok {
		if f.l.Len() == f.cap {
			delete(f.m, f.l.Back().Value.(data).key)
			f.l.Remove(f.l.Back())
		}
		el := f.l.PushFront(d)
		f.m[key] = el
	} else {
		el.Value = d
	}
	return true
}
