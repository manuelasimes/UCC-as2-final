package utils

import (
	"github.com/bradfitz/gomemcache/memcache"
	"fmt"
)

var cache *memcache.Client

func Init_cache() {
	cache = memcache.New("localhost:11211")
}

func Set(id int, startDate int, value []byte) {
    key := createCacheKey(id, startDate)
	cache.Set(&memcache.Item{Key: key, Value: value})
}


func Get(id int, startDate int) (value []byte) {
    key := createCacheKey(id, startDate)
    it, err := cache.Get(key)
    if err != nil {
        return nil
    }
    return it.Value
}

func createCacheKey(id int, startDate int) string {
    return fmt.Sprintf("reservation:%d:%d", id, startDate)
}