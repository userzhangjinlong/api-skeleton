package KafkaConsumer

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"sync"
)

var wg sync.WaitGroup

type Consumer struct {
	Topic     string
	Address   []string
	Partition int32
	cfg       *sarama.Config
}

func NewKafkaConsumer(topic string, addr []string, partition int32) *Consumer {
	return &Consumer{
		Topic:     topic,
		Address:   addr,
		Partition: partition,
	}
}

//NewKafkaConfig 初始化kafka配置
func (c *Consumer) NewKafkaConfig() {
	c.cfg = sarama.NewConfig()
	//下面可以自定义配置自己的kafka配置
	return
}

func (c *Consumer) PartitionConsumer(consumer sarama.Consumer) sarama.PartitionConsumer {
	partitionConsumer, err := consumer.ConsumePartition(c.Topic, c.Partition, sarama.OffsetNewest)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"topic":     c.Topic,
			"partition": c.Partition,
		}).Errorf("kafka PartitionConsumer err:%s", err)
		return nil
	}
	defer partitionConsumer.Close()

	return partitionConsumer
}
