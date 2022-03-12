package Util

import (
	"api-skeleton/app/Global"
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"time"
)

/**
*nsq 相关消息生产封装
*Publish()：阻塞发布1条消息。底层调用"PUB"指令
*PublishAsync()：非阻塞发布1条消息。相比Publish()，多了一个额外的doneChan参数，通过此chan来异步接收发布结果。
*MultiPublish()：阻塞发布多条消息。底层调用"MPUB"指令
*MultiPublishAsync()：非阻塞发布多条消息。通过doneChan来异步接收发布结果。
*DeferredPublish()：阻塞发布1条带延时的消息。相比Publish()，多了一个delay参数来指定延时多久才推送给消费者。底层调用"DPUB"指令
*DeferredPublishAsync()：非阻塞发布1条带延时的消息。通过doneChan来异步接收发布结果。
 */

func DeliveryNsq1Message(topic string, message []byte) {
	err := Global.NsqProducer.Publish(topic, message) // 发布消息
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":     err,
			"topic":   topic,
			"message": message,
		}).Error(fmt.Sprintf("producer.Publish,err : %v", err))
	}
}

func DeliveryDelayNsq1Message(topic string, message []byte, delayTime time.Duration) {
	err := Global.NsqProducer.DeferredPublish(topic, delayTime, message)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":       err,
			"topic":     topic,
			"delayTime": delayTime,
			"message":   message,
		}).Error(fmt.Sprintf("producer.delayPublish,err : %v", err))
	}
}

func DeliveryNsq2Message(topic string, message []byte) {
	err := Global.NsqProducer2.Publish(topic, message) // 发布消息
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":     err,
			"topic":   topic,
			"message": message,
		}).Error(fmt.Sprintf("producer.Publish,err : %v", err))
	}
}

func DeliveryNsq3Message(topic string, message []byte) {
	err := Global.NsqProducer3.Publish(topic, message) // 发布消息
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":     err,
			"topic":   topic,
			"message": message,
		}).Error(fmt.Sprintf("producer.Publish,err : %v", err))
	}
}

func CreateNsqConsumer(topic, chanel string, cfg *nsq.Config) (consumer *nsq.Consumer, err error) {
	consumer, err = nsq.NewConsumer(topic, chanel, cfg)
	if err != nil {
		fmt.Printf("异常:{\n %s \n}\n", err.Error())
		return nil, err
	}
	return
}
