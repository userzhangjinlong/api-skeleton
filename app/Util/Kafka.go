package Util

import (
	"api-skeleton/app/Global"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"log"
)

//SendKafkaProducerMsg kafka消息投递
func SendKafkaProducerMsg(topic string, message string) (pid int32) {
	kafkaCfg := sarama.NewConfig()
	// 设置生产者 消息 回复等级 0 1 all
	kafkaCfg.Producer.RequiredAcks = sarama.WaitForAll

	// 设置生产者 发送的分区
	kafkaCfg.Producer.Partitioner = sarama.NewRandomPartitioner

	//设置生产者 成功 发送消息 将在什么 通道返回
	kafkaCfg.Producer.Return.Successes = true

	// 新建一个同步生产者
	client, err := sarama.NewSyncProducer([]string{
		Global.Configs.Kafka.Node1,
	}, kafkaCfg)

	if err != nil {
		logrus.Errorf("kafkaProducer connect, err:", err)
		return
	}
	defer client.Close()
	// 定义一个生产消息，包括Topic、消息内容、
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	//msg.Key = sarama.StringEncoder(key)
	msg.Value = sarama.StringEncoder(message)

	// 发送消息
	pid, _, err = client.SendMessage(msg)

	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	return
}

//KafkaConsumerClient 初始化kafka消息消费者客户端
func KafkaConsumerClient(hosts []string, conf *sarama.Config) sarama.Consumer {
	client, err := sarama.NewClient(hosts, conf)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
	}
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer consumer.Close()

	return consumer
}
