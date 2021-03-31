package cache

import (
	"github.com/go-redis/redis/v8"
)

type ICacheEngine interface {
	GetCacheClient() *redis.Client
	CheckConnection() error
}

