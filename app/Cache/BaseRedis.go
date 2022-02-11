package Cache

import (
	"api-skeleton/app/Global"
	"github.com/garyburd/redigo/redis"
)

type RedisClient interface {
	SelectDB(db int)
	Get(key string) (string, error)
}

type BaseRedis struct {
}

//Get redis get
func (r *BaseRedis) Get(key string) (string, error) {
	val, err := redis.String(Global.RedisClient.Do("get", key))
	if err != nil {
		return "", err
	}

	return val, nil
}

//SelectDB redis选择库
func (r *BaseRedis) SelectDB(db int) {
	Global.RedisClient.Do("SELECT", db)
}
