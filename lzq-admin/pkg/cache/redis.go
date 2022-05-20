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
	token "lzq-admin/pkg/auth"
	"lzq-admin/pkg/hsflogger"
	"lzq-admin/pkg/utility"
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

type RedisHelper struct {
	isUseMultiTenancy bool
	prefixKey         string
}

func (r *RedisHelper) normalizeKey(key string) string {
	if len(key) == 0 {
		hsflogger.LogError("key不能为空", nil)
		panic(errors.New("key不能为空"))
	}
	nKey := key
	// 拼接缓存名称
	if r != nil && len(r.prefixKey) > 0 {
		nKey = fmt.Sprintf("%v:%v", r.prefixKey, nKey)
	}
	// 租户ID
	if r != nil && r.isUseMultiTenancy && config.Config.ServerConfig.UseMultiTenancy && len(token.GlobalTokenClaims.TenantId) > 0 {
		nKey = fmt.Sprintf("t:%v:%v", token.GlobalTokenClaims.TenantId, nKey)
	} else {
		nKey = fmt.Sprintf("c:%v", nKey)
	}
	// 缓存池
	if len(config.Config.RedisConfig.RedisPoolName) > 0 {
		nKey = fmt.Sprintf("%v:%v", config.Config.RedisConfig.RedisPoolName, nKey)
	}
	return nKey
}

func GetDefaultExpiresAt(expiration time.Duration) time.Duration {
	if expiration <= 0 {
		return time.Duration(24*60*60+utility.RandomNum(0, 61)) * time.Second
	} else {
		return expiration
	}
}

func (r *RedisHelper) Get(key string) string {
	val, err := redisClient.Get(ctx, r.normalizeKey(key)).Result()
	if err != nil {
		return ""
	}
	return val
}
func (r *RedisHelper) Set(key string, value interface{}, expiration time.Duration) {
	if err := redisClient.Set(ctx, r.normalizeKey(key), value, GetDefaultExpiresAt(expiration)).Err(); err != nil {
		panic(err)
	}
}
func (r *RedisHelper) SSet(key string, value interface{}, expiration time.Duration) {
	json, _ := jsoniter.MarshalToString(value)
	r.Set(key, json, expiration)
}
func (r *RedisHelper) Delete(key string) {
	redisClient.Del(ctx, r.normalizeKey(key))
}
func (r *RedisHelper) Keys(pattern string) []string {
	val, err := redisClient.Keys(ctx, pattern).Result()
	if err != nil {
		return []string{}
	}
	return val
}
func (r *RedisHelper) MultiGet(keys []string) []interface{} {
	keyNs := make([]string, 0)
	for _, v := range keys {
		keyNs = append(keyNs, r.normalizeKey(v))
	}
	val, err := redisClient.MGet(ctx, keyNs...).Result()
	if err != nil {
		return []interface{}{}
	}
	return val
}
func (r *RedisHelper) MultiDelete(keys []string) {
	keyNs := make([]string, 0)
	for _, v := range keys {
		keyNs = append(keyNs, r.normalizeKey(v))
	}
	redisClient.Del(ctx, keyNs...)
}

func (r *RedisHelper) HSet(key, field string, value interface{}) {
	if err := redisClient.HSet(ctx, r.normalizeKey(key), field, value).Err(); err != nil {
		panic(err)
	}
}
func (r *RedisHelper) HGet(key, field string) interface{} {
	val, err := redisClient.HGet(ctx, r.normalizeKey(key), field).Result()
	if err != nil {
		return ""
	}
	return val
}
func (r *RedisHelper) HDelete(key string, fields ...string) interface{} {
	val, err := redisClient.HDel(ctx, r.normalizeKey(key), fields...).Result()
	if err != nil {
		return ""
	}
	return val
}
func (r *RedisHelper) HGetAll(key string) interface{} {
	val, err := redisClient.HGetAll(ctx, r.normalizeKey(key)).Result()
	if err != nil {
		return ""
	}
	return val
}

func (r *RedisHelper) IncrBy(key string, increment int64) int64 {
	val, err := redisClient.IncrBy(ctx, r.normalizeKey(key), increment).Result()
	if err != nil {
		panic(err)
	}
	return val
}
