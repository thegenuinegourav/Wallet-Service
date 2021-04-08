package cache

import (
	"context"
	"github.com/WalletService/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// distributed cron caching to make sure cron gets executes only for one instance at same time by locking mechanism
type cronCache struct {
	expires time.Duration
}

var (
	cCCE  ICacheEngine
	cCCtx context.Context
)

type ICronCache interface {
	AcquireLock(key string, value string) bool
	ReleaseLock(key string, client string)
}

func NewCronCache(cacheEngine ICacheEngine, config config.PropertyM) ICronCache {
	cCCE = cacheEngine
	cCCtx = context.Background()
	return &cronCache{time.Duration(config.Expiry) * time.Minute}
}

func (cC *cronCache) AcquireLock(key string, value string) bool {
	check, err := cCCE.GetCacheClient().SetNX(cCCtx, key, value, cC.expires).Result()
	if err != nil {
		log.Println("Error occurred while setting cron cache, error : ", err)
	}
	log.Printf("Cron Cache | AcquireLock Key : %s, Value : %s, Check : %v \n", key, value, check)
	return check
}

func (cC *cronCache) ReleaseLock(key string, client string) {
	val, err := cCCE.GetCacheClient().Get(cCCtx, key).Result()
	if err != redis.Nil && val == client {
		cCCE.GetCacheClient().Del(cCCtx, key)
	}else {
		log.Printf("Error occurred while releasing lock for key %s , error : %v", key, err)
	}
	log.Printf("Cron Cache | ReleaseLock Key : %s, Client : %s \n", key, client)
}