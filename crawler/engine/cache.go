package engine

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

var cachePre = "go:cache:url"

type CacheDriver struct {
	redis *redis.Client
}

func CreateCacheDriver() CacheDriver {
	return CacheDriver{
		redis: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})}
}

func (c CacheDriver) IsDuplicate(v string) bool  {
	h := md5.New()
	h.Write([]byte(v))
	str := hex.EncodeToString(h.Sum(nil))
	key := cachePre + str
	val := c.redis.Exists(context.Background(), key).Val()
	if val > 0 {
		return true
	}

	value := c.redis.Set(context.Background(), key,1, 3600 * time.Second).Val()
	if value != "OK" {
		panic(errors.New("cache set failed :" + v))
	}
	return false
}
