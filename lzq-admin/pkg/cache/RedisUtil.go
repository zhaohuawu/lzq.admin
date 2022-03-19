package cache

/**
 * @Author  糊涂的老知青
 * @Date    2022/3/11
 * @Version 1.0.0
 */

var RedisUtil = &redisUtil{useMultiTenancy: true}

func UseMultiTenancy(useMultiTenancy bool) *redisUtil {
	r := &redisUtil{useMultiTenancy: useMultiTenancy}
	return r
}
