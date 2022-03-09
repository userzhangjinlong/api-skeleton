package bootstrap

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Global"
	"api-skeleton/app/Util"
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	"fmt"
	"github.com/nsqio/go-nsq"
)

//InitConfig 初始化config配置
func InitConfig() {
	Global.Configs = configs
}

//InitDB 初始化mysq
func InitDB() {
	db, err := ConnectPoolFactory.GetMysql(ConstDir.DEFAULT)
	if err != nil {
		panic("db链接获取异常")
	}
	Global.DB = db
}

//InitRedisClient 初始化redis
func InitRedisClient() {
	//初始化设置全局变量
	redisPool, _ := ConnectPoolFactory.GetRedis()
	Global.RedisClient = redisPool
}

//InitRedisClusterClient 初始化redis-cluster
func InitRedisClusterClient() {
	//初始化设置全局变量
	redisPool, _ := ConnectPoolFactory.GetRedisCluster()
	Global.RedisCluster = redisPool
}

//InitTracer 初始化jaegerTracer
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

//InitNsqProducer 初始化nsq生产者
func InitNsqProducer() {
	nsqAddress := fmt.Sprintf("%s:%s", configs.Nsq.Host, configs.Nsq.Node1)
	nsqConfig := nsq.NewConfig()
	nsqConfig.AuthSecret = configs.Nsq.Password
	producer, err := nsq.NewProducer(nsqAddress, nsqConfig)
	if err != nil {
		panic(err)
	}

	err = producer.Ping()
	if err != nil {
		//关闭生产者
		producer.Stop()
		panic("nsq链接异常")
	}

	Global.NsqProducer = producer
}
