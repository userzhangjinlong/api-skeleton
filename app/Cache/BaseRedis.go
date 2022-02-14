package Cache

import (
	"api-skeleton/app/Global"
	"github.com/garyburd/redigo/redis"
)

type RedisClient interface {
	SelectDB(db int)
	Expire(key string, ttl int) (bool, error)
	Get(key string) (string, error)
	Set(key string, value interface{}) (bool, error)
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
	val, err := redis.Bool(Global.RedisClient.Do("expire", key, ttl))
	if err != nil {
		return val, err
	}

	return val, nil
}

//Get redis get
func (r *BaseRedis) Get(key string) (string, error) {
	val, err := redis.String(Global.RedisClient.Do("get", key))
	if err != nil {
		return "", err
	}

	return val, nil
}

//Set set key value
func (r *BaseRedis) Set(key string, value interface{}) (bool, error) {
	_, err := Global.RedisClient.Do("set", key, value)
	if err != nil {
		return false, err
	}

	return true, nil
}

//HSet redis hash set
func (r *BaseRedis) HSet(key string, field string, value interface{}) (bool, error) {
	val, err := redis.Bool(Global.RedisClient.Do("hSet", key, field, value))
	if err != nil {
		return val, err
	}

	return val, nil
}

//HGet redis hash get string
func (r *BaseRedis) HGet(key string, name string) (string, error) {
	val, err := redis.String(Global.RedisClient.Do("hGet", key, name))
	if err != nil {
		return "", err
	}

	return val, nil
}

//LPush redis list Lpush
func (r *BaseRedis) LPush(key string, value interface{}) (bool, error) {
	val, err := redis.Bool(Global.RedisClient.Do("lpush", key, value))
	if err != nil {
		return val, err
	}

	return val, nil
}

//LPop redis list lpop
func (r *BaseRedis) LPop(key string) error {
	_, err := redis.Bool(Global.RedisClient.Do("lpop", key))
	if err != nil {
		return err
	}

	return nil
}
