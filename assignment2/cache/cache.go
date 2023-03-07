package cache

type Cache interface {
	Get(key string) any
	Set(key string, value any) bool
}
