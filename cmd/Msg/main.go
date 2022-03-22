package main

import (
	"api-skeleton/app/Consumer"
	"flag"
)

var (
	msgType = flag.String("MT", "", "消息类型")
	topic   = flag.String("T", "", "消息topic")
	chanel  = flag.String("C", "", "消息chanel")
)

func main() {
	flag.Parse()
	if *topic == "" {
		panic("请传入T参数 消息topic")
	}
	if *chanel == "" {
		panic("请传入C参数 消息chanel")
	}
	if *msgType == "" {
		*msgType = "nsq"
	}
	consumerInit := Consumer.NewMsgType(*msgType, *topic, *chanel)
	consumerInit.InitMsgConsumer()
	consumerInit.DistributionMsg()

}
