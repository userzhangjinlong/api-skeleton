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
	nsqAddress2 := fmt.Sprintf("%s:%s", configs.Nsq.Host, configs.Nsq.Node2)
	nsqAddress3 := fmt.Sprintf("%s:%s", configs.Nsq.Host, configs.Nsq.Node3)
	nsqConfig := nsq.NewConfig()
	nsqConfig.AuthSecret = configs.Nsq.Password
	producer, err := nsq.NewProducer(nsqAddress, nsqConfig)
	if err != nil {
		panic(err)
	}

	producer2, err := nsq.NewProducer(nsqAddress2, nsqConfig)
	if err != nil {
		panic(err)
	}

	producer3, err := nsq.NewProducer(nsqAddress3, nsqConfig)
	if err != nil {
		panic(err)
	}

	err = producer.Ping()
	if err != nil {
		//关闭生产者
		producer.Stop()
		panic("nsq节点1链接异常")
	}

	err = producer2.Ping()
	if err != nil {
		//关闭生产者
		producer2.Stop()
		panic("nsq节点2链接异常")
	}

	err = producer3.Ping()
	if err != nil {
		//关闭生产者
		producer3.Stop()
		panic("nsq节点3链接异常")
	}

	Global.NsqProducer = producer
	Global.NsqProducer2 = producer2
	Global.NsqProducer3 = producer3
}
