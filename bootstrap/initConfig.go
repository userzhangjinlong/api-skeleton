package bootstrap

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Global"
	"api-skeleton/app/Util"
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	"fmt"
)

func InitDB() {
	db, err := ConnectPoolFactory.GetMysql(ConstDir.DEFAULT)
	if err != nil {
		panic("db链接获取异常")
	}
	Global.DB = db
}

func InitRedisClient() {
	//初始化设置全局变量
	redisPool, _ := ConnectPoolFactory.GetRedis()
	Global.RedisClient = redisPool
}

func InitRedisClusterClient() {
	//初始化设置全局变量
	redisPool, _ := ConnectPoolFactory.GetRedisCluster()
	Global.RedisCluster = redisPool
}

func InitTracer() {
	jaegerTracer, _, err := Util.NewJaegerTracer(
		configs.Trace.Servicename,
		fmt.Sprintf("%s%s", configs.Trace.Agenthost, configs.Trace.Port),
	)

	if err != nil {
		//todo::异常日志记录链接tracer链路追综异常
		return
	}
	Global.Tracer = jaegerTracer
	return
}

func InitConfig() {
	Global.Configs = configs
}
