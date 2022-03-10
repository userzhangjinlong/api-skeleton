package Util

import (
	"api-skeleton/app/Global"
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

func DeliveryNsqMessage(topic string, message []byte) {
	err := Global.NsqProducer.Publish(topic, message) // 发布消息
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
