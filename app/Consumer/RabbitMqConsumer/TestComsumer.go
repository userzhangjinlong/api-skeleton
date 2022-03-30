package RabbitMqConsumer

import (
	"api-skeleton/app/Libraries/RabbitMq"
	"fmt"
)

//TestConsumer 测试简易mq消息消费
func TestConsumer(queue string) {
	r := RabbitMq.NewRabbitMQ(queue, "testExchange", "testKey")
	msgs, err := r.GetSimpleConsumerMsg()
	if err != nil {
		panic(fmt.Sprintf("rabbitMq consumer err:%s", err))
	}
	handlefunc := func() {
		for d := range msgs {
			// 消息处理逻辑
			fmt.Printf("测试简易rabbitMq 消费消息：%v\n", string(d.Body))
		}
	}
	r.ConsumerSimpleMsg(handlefunc)
}
