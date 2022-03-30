package Consumer

import "api-skeleton/app/Consumer/RabbitMqConsumer"

func (m MsgType) DistributionMsg() {
	switch m.Cate {
	case "nsq":
		distributionNsqMsg(m.Topic)
	case "kafka":
		distributionKafkaMsg(m.Topic)
	case "rabbitMq":
		distributionRabbitMqMsg(m.Topic)
	default: //默认nsq消息
		distributionNsqMsg(m.Topic)
	}

}

//distributionNsqMsg 分发nsq消息consumer
func distributionNsqMsg(topic string) {
	switch topic {
	case "createRankingMessage":
		nsqConsumer.CreateRankingListConsumer()
	case "createRankingMessageNode2":
		nsqConsumer.CreateRankingListConsumer()
	case "createRankingMessageNode3":
		nsqConsumer.CreateRankingListConsumer()
	case "sendCoupon":
		nsqConsumer.SendCouponConsumer()
	}
}

func distributionKafkaMsg(topic string) {
	switch topic {
	case "kafka-test-1":
		kafkaConsumer.TestConsumer()
	default:
		kafkaConsumer.TestConsumer()
	}
}

func distributionRabbitMqMsg(topic string) {
	switch topic {
	case "testQueue":
		RabbitMqConsumer.TestConsumer(topic)
	default:
		RabbitMqConsumer.TestConsumer("testQueue")
	}
}
