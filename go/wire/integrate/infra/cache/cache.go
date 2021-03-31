package cache

import (
	"github.com/go-redis/cache/v8"
	"github.com/google/wire"

	"github.com/wenzong/demo/infra/redis"
)

func NewCache(r *redis.DefaultRedis) *cache.Cache {

	return cache.New(&cache.Options{
		Redis:        r,
		StatsEnabled: true,
	})
}

var ProviderSet = wire.NewSet(NewCache)
