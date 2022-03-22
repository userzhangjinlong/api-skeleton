package KafkaConsumer

import (
	"api-skeleton/app/Util"
	"fmt"
	"github.com/Shopify/sarama"
)

func (c *Consumer) TestConsumer() {
	c.NewKafkaConfig()
	kafkaConsumer := Util.KafkaConsumerClient(c.Address, c.cfg)
	partitionList, err := kafkaConsumer.Partitions(c.Topic)
	//这一步暂时待处理
	//partitionConsumer := c.PartitionConsumer(kafkaConsumer)
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)

	//遍历分区
	for partition := range partitionList {
		// 针对每个分区创建一个对应的分区消费者
		pc, err := kafkaConsumer.ConsumePartition(c.Topic, int32(partition), sarama.OffsetNewest)
		//pc, err := consumer.ConsumePartition("web_log", int32(partition), 90)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		// 异步从每个分区消费信息
		wg.Add(1) //+1
		go func(sarama.PartitionConsumer) {
			defer wg.Done() //-1
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait()
}
