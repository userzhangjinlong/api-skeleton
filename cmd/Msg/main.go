package main

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Model/ApiSkeleton"
	"api-skeleton/config"
	ConnectPoolFactory "api-skeleton/database/ConnectPool"
	"fmt"
	"github.com/nsqio/go-nsq"
	"math/rand"
	"strconv"
	"time"
)

type MessageHandler struct {
	nsqConsumer *nsq.Consumer
}

var configs = config.InitConfig

func main() {
	//定义消费者topic和nsq服务地址
	var (
		topic   = "createRankingMessage"
		address = fmt.Sprintf("%s:%s", configs.Nsq.Host, configs.Nsq.CusNode)
	)
	//实例化配置文件
	cfg := nsq.NewConfig()
	cfg.AuthSecret = configs.Nsq.Password
	cfg.LookupdPollInterval = 3 * time.Second
	//设置重连时间
	consumer, err := nsq.NewConsumer(topic, topic, cfg)
	if err != nil {
		fmt.Printf("异常:{\n %s \n}\n", err.Error())
	}
	handler := MessageHandler{nsqConsumer: consumer}
	consumer.AddHandler(&handler)
	err = consumer.ConnectToNSQLookupd(address)
	if err != nil {
		fmt.Printf("异常:{\n %s \n}\n", err.Error())
	}
	select {}
}

func (m MessageHandler) HandleMessage(message *nsq.Message) error {
	shop_id, _ := strconv.ParseInt(string(message.Body), 10, 64)
	createData := ApiSkeleton.ApiDoubleFlowBringRankRealtimeIncr1d{
		ShopId:          shop_id,
		FirstCategoryId: rand.Int63n(10),
		PopValue:        rand.Float64() + 60,
		Ranking:         rand.Int63n(100),
		ModifyTime:      time.Now().Format("2006-01-02 15:04:05"),
		Dt:              "2022-03-08",
	}
	//fmt.Println(createData)
	db, err := ConnectPoolFactory.GetMysql(ConstDir.DEFAULT)
	if err != nil {
		panic("db链接获取异常")
	}
	db.Create(&createData)
	return nil
}
