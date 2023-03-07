package lru

import (
	"container/list"
	"lru/cache"
	"lru/cache/redis"
)

type LRUCache struct {
	m              map[string]*list.Element
	cap            int
	l              list.List
	nextCacheLayer cache.Cache
}

func NewLRUCache(cap int) cache.Cache {

	return &LRUCache{
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

func (c *LRUCache) Get(key string) any {
	el, ok := c.m[key]
	if !ok {
		val := c.nextCacheLayer.Get(key)
		c.Set(key, val)
		return val
	}
	c.l.MoveToFront(el)
	return el.Value.(data).value
}

func (c *LRUCache) Set(key string, value any) bool {
	if c.cap <= 0 {
		return false
	}
	el, ok := c.m[key]
	d := data{key: key, value: value}
	if !ok {
		if c.l.Len() == c.cap {
			del := c.l.Back().Value.(data).key
			delete(c.m, del)
			c.l.Remove(c.l.Back())
		}
		el := c.l.PushFront(d)
		c.m[key] = el
	} else {
		el.Value = d
		c.m[key] = el
		c.l.MoveToFront(el)
	}
	return true
}
