package main

import (
	"api-skeleton/app/Consumer/NsqConsumer"
	"api-skeleton/bootstrap"
	"api-skeleton/config"
	"flag"
	"fmt"
)

var configs = config.InitConfig

var (
	nsqConsumer *NsqConsumer.Consumer
	topic       = flag.String("T", "", "消息topic")
	chanel      = flag.String("C", "", "消息chanel")
	address     = fmt.Sprintf("%s:%s", configs.Nsq.Host, configs.Nsq.CusNode)
)

func main() {
	flag.Parse()
	if *topic == "" {
		panic("请传入T参数 消息topic")
	}
	if *chanel == "" {
		panic("请传入C参数 消息chanel")
	}
	//fmt.Println(*topic)
	//fmt.Println(*chanel)
	//初始化配置
	bootstrap.InitConfig()
	nsqConsumer = new(NsqConsumer.Consumer)
	nsqConsumer.Address = address
	nsqConsumer.Topic = *topic
	nsqConsumer.Chanel = *chanel
	//设置重连时间
	switch *topic {
	case "createRankingMessage":
		nsqConsumer.CreateRankingListConsumer()
	case "sendCoupon":
		nsqConsumer.SendCouponConsumer()
	}

}
