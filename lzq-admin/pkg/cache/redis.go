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
	jsoniter "github.com/json-iterator/go"
	"lzq-admin/config"
	"lzq-admin/middleware"
	"lzq-admin/pkg/hsflogger"
	"math/rand"
	"time"
)

// RedisClient Redis缓存客户端单例
var redisClient *redis.Client
var ctx = context.Background()

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisConfig.RedisHost,
		Password: config.Config.RedisConfig.RedisPwd,
		DB:       config.Config.RedisConfig.RedisDB,
	})
	if _, err := client.Ping(ctx).Result(); err != nil {
		hsflogger.LogError("", err)
		panic(err)
	}
	redisClient = client
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
	if config.Config.ServerConfig.UseMultiTenancy && len(middleware.TokenClaims.TenantId) > 0 {
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
		return expiration
	}
}

func (r *redisUtil) Get(key string) string {
	val, err := redisClient.Get(ctx, normalizeKey(key)).Result()
	if err != nil {
		return ""
	}
	return val
}

func (r *redisUtil) SetEx(key string, value string, expiration time.Duration) {
	err := redisClient.SetEX(ctx, normalizeKey(key), value, GetDefaultExpiresAt(expiration)).Err()
	if err != nil {
		panic(err)
	}
}

func (r *redisUtil) Delete(key string) {
	redisClient.Del(ctx, normalizeKey(key))
}

func (r *redisUtil) Keys(pattern string) []string {
	val, err := redisClient.Keys(ctx, pattern).Result()
	if err != nil {
		return []string{}
	}
	return val
}

func (r *redisUtil) MultiGet(keys []string) []interface{} {
	keyNs := make([]string, 0)
	for _, v := range keys {
		keyNs = append(keyNs, normalizeKey(v))
	}
	val, err := redisClient.MGet(ctx, keyNs...).Result()
	if err != nil {
		return []interface{}{}
	}
	return val
}

func (r *redisUtil) MultiDelete(keys []string) {
	keyNs := make([]string, 0)
	for _, v := range keys {
		keyNs = append(keyNs, normalizeKey(v))
	}
	redisClient.Del(ctx, keyNs...)
}

func (r *redisUtil) SetString(key string, value string, expiration time.Duration) {
	if err := redisClient.Set(ctx, normalizeKey(key), value, GetDefaultExpiresAt(expiration)).Err(); err != nil {
		panic(err)
	}
}
func (r *redisUtil) SetInterface(key string, value interface{}, expiration time.Duration) {
	defaultExpiresAt := GetDefaultExpiresAt(expiration)
	fmt.Println("defaultExpiresAt：" + defaultExpiresAt.String())
	if err := redisClient.Set(ctx, normalizeKey(key), value, defaultExpiresAt).Err(); err != nil {
		panic(err)
	}
}
func (r *redisUtil) SetInterfaceArray(key string, value interface{}, expiration time.Duration) {
	jsonString, _ := jsoniter.MarshalToString(value)
	r.SetString(key, jsonString, expiration)
}

func (r *redisUtil) IncrBy(key string, increment int64) int64 {
	val, err := redisClient.IncrBy(ctx, normalizeKey(key), increment).Result()
	if err != nil {
		panic(err)
	}
	return val
}
