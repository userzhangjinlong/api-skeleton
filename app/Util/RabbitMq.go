package Util

import "api-skeleton/app/Libraries/RabbitMq"

//SendRabbitMqMsg 统一发送mq消息服务入口 测试mq消息 topic channel exchange queue 这四个mq的东西需要后期确认详细使用
func SendRabbitMqMsg(queue, exchange, msg string) {
	rabbitMq := RabbitMq.NewRabbitMQ(queue, exchange, "testKey")
	rabbitMq.ProducerMsg(msg)
}
