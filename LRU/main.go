package main

import (
	"fmt"
	"lru/cache/lru"
)

func main() {
	cache := lru.NewLRUCache(2)
	cache.Set("india", "delhi")
	cache.Set("pakistan", "islamabad")
	fmt.Println(cache.Get("india"))
	cache.Set("bangladesh", "dhaka")
	fmt.Println(cache.Get("india"))
	fmt.Println(cache.Get("Nepal"))
}
