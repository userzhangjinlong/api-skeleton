package Util

import (
	"api-skeleton/app/Global"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func SendKafkaProducerMsg(topic string, key string, message string) (pid int32) {
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
		//Global.Configs.Kafka.Node2,
		//Global.Configs.Kafka.Node3,
	}, kafkaCfg)

	if err != nil {
		logrus.Errorf("kafkaProducer connect, err:", err)
		return
	}
	defer client.Close()
	fmt.Println(key)
	// 定义一个生产消息，包括Topic、消息内容、
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	//msg.Key = sarama.StringEncoder(key)
	msg.Value = sarama.StringEncoder(message)

	//发送消息前判断topic是否存在 并创建

	// 发送消息
	pid, _, err = client.SendMessage(msg)

	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	return
}
