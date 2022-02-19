package cache

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/29
 * @Version 1.0.0
 */

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"lzq-admin/config"
	"lzq-admin/middleware"
	"lzq-admin/pkg/hsflogger"
	"math/rand"
	"time"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client
var ctx = context.Background()

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisConfig.RedisHost,
		Password: config.Config.RedisConfig.RedisPwd,
		DB:       config.Config.RedisConfig.RedisDB,
	})
	if _, err := client.Ping(ctx).Result(); err != nil {
		hsflogger.LogError("", err)
		panic(err)
	}
	RedisClient = client
}

type redisUtil struct{}

var RedisUtil = &redisUtil{}

func normalizeKey(key string) string {
	if len(key) == 0 {
		hsflogger.LogError("key不能为空", nil)
		panic(errors.New("key不能为空"))
	}
	nKey := key
	// 租户ID
	if len(middleware.TokenClaims.TenantId) > 0 {
		nKey = fmt.Sprintf("%v:%v", middleware.TokenClaims.TenantId, nKey)
	}
	// 缓存池
	if len(config.Config.RedisConfig.RedisPoolName) > 0 {
		nKey = fmt.Sprintf("%v:%v", config.Config.RedisConfig.RedisPoolName, nKey)
	}
	return nKey
}

func GetDefaultExpiresAt(expiration time.Duration) time.Duration {
	if expiration <= 0 {
		return time.Duration(24*60*60+rand.Intn(60)) * time.Second
	} else {
		return expiration * time.Second
	}
}

func (*redisUtil) Get(key string) string {
	val, err := RedisClient.Get(ctx, normalizeKey(key)).Result()
	if err != nil {
		return ""
	}
	return val
}

func (*redisUtil) SetEx(key string, value string, expiration time.Duration) {
	err := RedisClient.SetEX(ctx, normalizeKey(key), value, GetDefaultExpiresAt(expiration)).Err()
	if err != nil {
		panic(err)
	}
}

func (*redisUtil) Delete(key string) {
	RedisClient.Del(ctx, normalizeKey(key))
}

func (*redisUtil) Keys(pattern string) []string {
	val, err := RedisClient.Keys(ctx, pattern).Result()
	if err != nil {
		return []string{}
	}
	return val
}

func (*redisUtil) MultiGet(keys []string) []interface{} {
	keyNs := make([]string, 0)
	for _, v := range keys {
		keyNs = append(keyNs, normalizeKey(v))
	}
	val, err := RedisClient.MGet(ctx, keyNs...).Result()
	if err != nil {
		return []interface{}{}
	}
	return val
}

func (*redisUtil) MultiDelete(keys []string) {
	keyNs := make([]string, 0)
	for _, v := range keys {
		keyNs = append(keyNs, normalizeKey(v))
	}
	RedisClient.Del(ctx, keyNs...)
}

func (*redisUtil) Set(key string, value string, expiration time.Duration) {
	if err := RedisClient.Set(ctx, normalizeKey(key), value, GetDefaultExpiresAt(expiration)).Err(); err != nil {
		panic(err)
	}
}

func (*redisUtil) IncrBy(key string, increment int64) int64 {
	val, err := RedisClient.IncrBy(ctx, normalizeKey(key), increment).Result()
	if err != nil {
		panic(err)
	}
	return val
}
