package Cache

import (
	"api-skeleton/app/Global"
	"time"
)

type RedisClient interface {
	SelectDB(db int)
	Expire(key string, ttl int) (bool, error)
	Get(key string) (string, error)
	Set(key string, value interface{}, expireTime time.Duration) (bool, error)
	HSet(key string, field string, value interface{}) (bool, error)
	HGet(key string, name string) (string, error)
}

type BaseRedis struct {
}

//SelectDB redis选择库
func (r *BaseRedis) SelectDB(db int) {
	Global.RedisClient.Do("SELECT", db)
}

//Expire redis key expire
func (r *BaseRedis) Expire(key string, ttl int) (bool, error) {
	val, err := Global.RedisClient.Expire(key, time.Duration(ttl)).Result()
	if err != nil {
		return val, err
	}

	return val, nil
}

//Get redis get
func (r *BaseRedis) Get(key string) (string, error) {
	val, err := Global.RedisClient.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

//Set set key value
func (r *BaseRedis) Set(key string, value interface{}, ttl time.Duration) (bool, error) {
	_, err := Global.RedisClient.Set(key, value, ttl).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

//HSet redis hash set
func (r *BaseRedis) HSet(key string, field string, value interface{}) (bool, error) {
	val, err := Global.RedisClient.HSet(key, field, value).Result()
	if err != nil {
		return val, err
	}

	return val, nil
}

//HGet redis hash get string
func (r *BaseRedis) HGet(key string, name string) (string, error) {
	val, err := Global.RedisClient.HGet(key, name).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

//LPush redis list Lpush
func (r *BaseRedis) LPush(key string, value interface{}) (int64, error) {
	val, err := Global.RedisClient.LPush(key, value).Result()
	if err != nil {
		return val, err
	}

	return val, nil
}

//LPop redis list lpop
func (r *BaseRedis) LPop(key string) error {
	_, err := Global.RedisClient.LPop(key).Result()
	if err != nil {
		return err
	}

	return nil
}
