package ConnectPoolFactory

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
	"web_go/Utils/Config"
)

var configs Config.Config

//redis工厂加载redis连接池
func NewRedis(db ...int) (result bool) {
	for _, v := range db {
		if v != 0 {
			redisDb = v
		} else {
			redisDb, _ = strconv.Atoi(configs.GetInstance().GetString("db"))
		}
	}
	result = NewConnect("redis").GetInstance().InitConnectPool()

	return result
}

//redis工厂选择加载的redis库
func SelectDb(db int) (result bool) {
	redisDb = db
	result = NewConnect("redis").GetInstance().InitConnectPool()

	return result
}

//redis工厂获取redis连接池
func GetRedis() (redisPool *redis.Pool, err error) {
	redisConnect, errRedis := NewConnect("redis").GetInstance().GetConnectLibrary()

	return redisConnect.(*redis.Pool), errRedis
}
