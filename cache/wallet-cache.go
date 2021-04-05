package cache

import (
	"encoding/json"
	"github.com/WalletService/config"
	"github.com/WalletService/model"
	"github.com/go-redis/redis/v8"
	"log"
	"context"
	"strconv"
	"time"
)

type walletCache struct {
	expires time.Duration
}

var (
	cE ICacheEngine
	ctx context.Context
)

type IWalletCache interface {
	Set(key int, value *model.Wallet) error
	Get(key int) *model.Wallet
}

func NewWalletCache(cacheEngine ICacheEngine, config config.Property) IWalletCache {
	cE = cacheEngine
	ctx = context.Background()
	return &walletCache{time.Duration(config.Expiry) * time.Hour}
}

func (wC *walletCache) Set(key int, value *model.Wallet) error {
	wallet, err := json.Marshal(value)
	if err != nil {
		log.Println("Error occurred while marshalling cache wallet value to model")
		return err
	}
	err = cE.GetCacheClient().Set(ctx, strconv.Itoa(key), wallet, wC.expires).Err()
	if err != nil {
		log.Println("Error occurred while pushing Wallet to redis, error : ", err)
		return err
	}
	log.Printf("Cached | Key : %d , Value : %v", key, value)
	return nil
}

func (wC *walletCache) Get(key int) *model.Wallet {
	val, err := cE.GetCacheClient().Get(ctx, strconv.Itoa(key)).Result()
	if err == redis.Nil {
		log.Printf("Wallet ID not found : %d\n", key)
		return nil
	}else if err != nil {
		log.Println("Error occurred while fetching Wallet from redis, error : ", err)
		return nil
	}
	wallet := model.Wallet{}
	err = json.Unmarshal([]byte(val), &wallet)
	if err != nil {
		log.Println("Error occurred while unmarshalling Wallet model, error : ", err)
		return nil
	}
	log.Printf("Cached Result | Key : %d , Value : %v", key, wallet)
	return &wallet
}
