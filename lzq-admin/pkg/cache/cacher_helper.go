package cache

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/25
 * @Version 1.0.0
 */

type cacheHelper struct{}

var LzqCacheHelper = &cacheHelper{}

// GetCacheVersion 获取缓存类型版本号
func (c *cacheHelper) GetCacheVersion(cacheType string) string {
	key := c.GetCacheTypeVersionKey(cacheType)
	value := RedisUtil.NewRedis(false).Get(key)
	if len(value) == 0 {
		value = uuid.NewV4().String()
		RedisUtil.NewRedis(false).SetString(key, value, time.Hour*24)
	}
	return value
}
func (c *cacheHelper) GetCacheTypeVersionKey(cacheType string) string {
	return fmt.Sprintf("CacheVersion:%v", cacheType)
}
