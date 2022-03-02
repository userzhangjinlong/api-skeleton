package ConnectPoolFactory

import (
	"api-skeleton/config"
	"github.com/go-redis/redis"
	"strconv"
)

var configs = config.InitConfig

//NewRedis redis工厂加载redis连接池
func NewRedis(db ...int) (result bool) {
	for _, v := range db {
		if v != 0 {
			redisDb = v
		} else {
			redisDb, _ = strconv.Atoi(configs.Redis.Db)
		}
	}
	result = NewConnect("redis", "").GetInstance().InitConnectPool()

	return result
}

//SelectDb redis工厂选择加载的redis库
func SelectDb(db int) (result bool) {
	redisDb = db
	result = NewConnect("redis", "").GetInstance().InitConnectPool()

	return result
}

//GetRedis redis工厂获取redis连接池
func GetRedis() (redisPool *redis.Client, err error) {
	if !NewRedis() {
		panic("redis链接异常")
	}
	redisConnect, errRedis := NewConnect("redis", "").GetInstance().GetConnectLibrary()

	return redisConnect.(*redis.Client), errRedis
}
