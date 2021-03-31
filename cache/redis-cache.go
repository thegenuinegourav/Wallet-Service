package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

type redisCache struct {
	host string
	password string
	db int
}

var (
	rC 			*redisCache
	rCOnce 		sync.Once
	rClient 	*redis.Client
	rOnce 		sync.Once
)

// Making muxRouter instance as singleton
func NewCache(host string, password string, db int) ICacheEngine {
	if rC == nil {
		rCOnce.Do(func() {
			rC = &redisCache{
				host: host, password: password, db : db,
			}
		})
	}
	return rC
}

func InitRedisClient(r *redisCache) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.host,
		Password: r.password,
		DB:       r.db,
	})
	rClient = rdb
}

// Making sure redisCache only initialise once as singleton
func (r *redisCache) GetCacheClient() *redis.Client {
	if rClient == nil {
		rOnce.Do(func() {
			InitRedisClient(r)
		})
	}
	return rClient
}

func (r *redisCache) CheckConnection() error {
	pong, err := rClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Redis connection failed, error : ", err)
		return err
	}
	fmt.Println(pong, err)
	log.Println("Redis connection running on port 6379")
	// Output: PONG <nil>
	return nil
}
