package models

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis(host string) {
	rdb = redis.NewClient(&redis.Options{Addr: host})
}

func CacheGoodsToRedis(key string, value []byte, expiration time.Duration)error{
	return rdb.Set(ctx, key, value, expiration).Err()
}
func GetGoodsFromRedis(key string)(string,error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return val, nil
}
