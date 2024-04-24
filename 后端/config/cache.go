package config

import (
	"github.com/redis/go-redis/v9"
	Cache "github.com/wike2019/wike_go/lib/cache"
	CacheMemory "github.com/wike2019/wike_go/lib/cache/memory"
	CacheRedis "github.com/wike2019/wike_go/lib/cache/redis"
)

// 配置缓存模块 使用redis
func RedisCache(r *redis.Client) Cache.Cache {
	c := CacheRedis.NewCache(r)
	return c
}

// 配置缓存模块  使用内存
func MemoryCache() Cache.Cache {
	c := CacheMemory.NewCache()
	return c
}
