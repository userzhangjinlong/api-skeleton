package di

import (
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	"github.com/garyburd/redigo/redis"
)

func InitRedis() *redis.Pool {
	redis, err := ConnectPoolFactory.GetRedis()
	if err != nil {
		panic("redis链接获取异常")
	}

	return redis
}
