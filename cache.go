package gorm

import (
	"github.com/patrickmn/go-cache"
	"log"
)

func CacheSave(C *cache.Cache, key string, x interface{}){

	log.Println("[CACHE] Cache Saving: Key: %+v Object: %+v ", key, x)

	C.Set(key, x, cache.NoExpiration)
	valor, found := C.Get(key)

	log.Println("[CACHE] Cache Saved: Object: %+v Found: %+v ", valor, found)

}

func CacheGet(C *cache.Cache, key string) (interface{}, bool){

	valor, found := C.Get(key)
	log.Println("[CACHE] Cache GET: Object: %v Found: %v", valor, found)
	return valor, found

}


