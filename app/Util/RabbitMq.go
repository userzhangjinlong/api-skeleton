package Util

import "api-skeleton/app/Libraries/RabbitMq"

//SendRabbitMqMsg 统一发送mq消息服务入口
func SendRabbitMqMsg(queue, exchange, msg string) {
	rabbitMq := RabbitMq.NewRabbitMQ(queue, exchange, "testKey")
	rabbitMq.ProducerMsg(msg)
}
