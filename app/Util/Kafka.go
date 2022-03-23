package Util

import (
	"api-skeleton/app/Global"
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
	"time"
)

//SendKafkaProducerMsg kafka同步生产者消息投递
func SendKafkaProducerMsg(topic string, message string) (pid int32) {
	kafkaCfg := sarama.NewConfig()
	// 设置生产者 消息 回复等级 0 1 all 发送完数据需要leader和follow都确认
	kafkaCfg.Producer.RequiredAcks = sarama.WaitForAll

	// 设置生产者 发送的分区 随机分配分区 partition
	kafkaCfg.Producer.Partitioner = sarama.NewRandomPartitioner

	//设置生产者 成功 发送消息 将在什么 通道返回 成功交付的消息将在success channel返回
	kafkaCfg.Producer.Return.Successes = true
	//ack确认 一般 0对应异步 不需要确认只管消息处理速度 1对应同步
	//kafkaCfg.Producer.RequiredAcks = 1

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

//SendKafkaSyncProducerMsg kafka异步消息投递
func SendKafkaSyncProducerMsg(topic string, message string) {
	kafkaCfg := sarama.NewConfig()
	// 异步生产者不建议把 Errors 和 Successes 都开启，一般开启 Errors 就行
	// 同步生产者就必须都开启，因为会同步返回发送成功或者失败
	kafkaCfg.Producer.Return.Errors = true    // 设定是否需要返回错误信息
	kafkaCfg.Producer.Return.Successes = true // 设定是否需要返回成功信息
	producer, err := sarama.NewAsyncProducer([]string{Global.Configs.Kafka.Node1}, kafkaCfg)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	var (
		wg                                   sync.WaitGroup
		enqueued, timeout, successes, errors int
	)
	//默认发送一个消息
	limit := 1
	// [!important] 异步生产者发送后必须把返回值从 Errors 或者 Successes 中读出来 不然会阻塞 sarama 内部处理逻辑 导致只能发出去一条消息
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			// log.Printf("[Producer] Success: key:%v msg:%+v \n", s.Key, s.Value)
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for e := range producer.Errors() {
			log.Printf("[Producer] Errors：err:%v msg:%+v \n", e.Msg, e.Err)
			errors++
		}
	}()

	// 异步发送
	for i := 0; i < limit; i++ {
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)}
		// 异步发送只是写入内存了就返回了，并没有真正发送出去
		// sarama 库中用的是一个 channel 来接收，后台 goroutine 异步从该 channel 中取出消息并真正发送
		// select + ctx 做超时控制,防止阻塞 producer.Input() <- msg 也可能会阻塞
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
		select {
		case producer.Input() <- msg:
			enqueued++
		case <-ctx.Done():
			timeout++
		}
		cancel()
		if i%10000 == 0 && i != 0 {
			log.Printf("已发送消息数:%d 超时数:%d\n", i, timeout)
		}
	}

	// We are done
	producer.AsyncClose()
	wg.Wait()
	log.Printf("发送完毕 总发送条数:%d enqueued:%d timeout:%d successes: %d errors: %d\n", limit, enqueued, timeout, successes, errors)
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
