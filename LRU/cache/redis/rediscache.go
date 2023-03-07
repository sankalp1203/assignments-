package redis

import "lru/cache"

type redisClient struct {
}

func NewRedis() cache.Cache {
	return &redisClient{}
}

func (r *redisClient) Get(key string) (d any) {
	return "generic answer"
}

func (r *redisClient) Set(key string, value any) bool {
	return true
}
