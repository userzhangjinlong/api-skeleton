package main

import (
	"api-skeleton/app/Consumer"
	"api-skeleton/app/Util"
	"flag"
)

var (
	msgType = flag.String("MT", "", "消息类型")
	topic   = flag.String("T", "", "消息topic")
	chanel  = flag.String("C", "", "消息chanel")
	msgArr  = []string{"kafka", "rabbitMq"}
)

func main() {
	flag.Parse()
	if *topic == "" {
		panic("请传入T参数 消息topic")
	}
	inArray := Util.InArray(*msgType, msgArr)
	if *chanel == "" && !inArray {
		panic("请传入C参数 消息chanel")
	}
	if *msgType == "" {
		*msgType = "nsq"
	}
	consumerInit := Consumer.NewMsgType(*msgType, *topic, *chanel)
	consumerInit.InitMsgConsumer()
	consumerInit.DistributionMsg()

}
