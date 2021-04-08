package cache

import (
	"encoding/json"
	"errors"
	"github.com/WalletService/config"
	"github.com/WalletService/model"
	"github.com/go-redis/redis/v8"
	"log"
	"context"
	"net/http"
	"time"
)

type transactionIdempotentCache struct {
	expires time.Duration
}

var (
	tICE ICacheEngine
	tICtx context.Context
)

type ITransactionIdempotentCache interface {
	GetIdempotencyKey(r *http.Request) (string, error)
	Set(key string, value *model.Transaction) error
	Get(key string) *model.Transaction
}

func NewTransactionIdempotentCache(cacheEngine ICacheEngine, config config.PropertyH) ITransactionIdempotentCache {
	tICE = cacheEngine
	tICtx = context.Background()
	return &transactionIdempotentCache{time.Duration(config.Expiry) * time.Hour}
}

func (tIC *transactionIdempotentCache) GetIdempotencyKey(r *http.Request) (string, error) {
	key := r.Header.Get("x-idempotency-key")
	if len(key) == 0 {
		return "", errors.New("x-idempotency-key should be present in headers for this request")
	}
	return key, nil
}

func (tIC *transactionIdempotentCache) Set(key string, value *model.Transaction) error {
	transaction, err := json.Marshal(value)
	if err != nil {
		log.Println("Error occurred while marshalling cache transaction value to model")
		return err
	}
	err = tICE.GetCacheClient().Set(tICtx, key, transaction, tIC.expires).Err()
	if err != nil {
		log.Println("Error occurred while pushing transaction to redis, error : ", err)
		return err
	}
	log.Printf("Cached | Key : %s , Value : %v", key, value)
	return nil
}

func (tIC *transactionIdempotentCache) Get(key string) *model.Transaction {
	val, err := tICE.GetCacheClient().Get(tICtx, key).Result()
	if err == redis.Nil {
		log.Printf("transaction ID not found : %s\n", key)
		return nil
	}else if err != nil {
		log.Println("Error occurred while fetching transaction from redis, error : ", err)
		return nil
	}
	transaction := model.Transaction{}
	err = json.Unmarshal([]byte(val), &transaction)
	if err != nil {
		log.Println("Error occurred while unmarshalling transaction model, error : ", err)
		return nil
	}
	log.Printf("Cached Result | Key : %s , Value : %v", key, transaction)
	return &transaction
}
