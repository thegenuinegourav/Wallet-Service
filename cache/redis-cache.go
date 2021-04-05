package cache

import (
	"context"
	"fmt"
	"github.com/WalletService/config"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

type redisCache struct {
	host     string
	password string
	db       int
	client   *redis.Client
	once     sync.Once
}

// Not making this singleton coz we might need to add different db server on redis
func NewCache(config config.Cache, db int) ICacheEngine {
	return &redisCache{
		host: config.Server + ":" + config.Port,
		password: config.Password,
		db : db,
	}
}

func InitRedisClient(r *redisCache) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.host,
		Password: r.password,
		DB:       r.db,
	})
	r.client = rdb
}

// Making sure redisCache only initialise once as singleton for single rediscache instance
func (r *redisCache) GetCacheClient() *redis.Client {
	if r.client == nil {
		r.once.Do(func() {
			InitRedisClient(r)
		})
	}
	return r.client
}

func (r *redisCache) CheckConnection() error {
	pong, err := r.client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Redis connection failed, error : ", err)
		return err
	}
	fmt.Println(pong, err)
	log.Printf("Redis DB %d connection running on port 6379\n", r.db)
	// Output: PONG <nil>
	return nil
}
