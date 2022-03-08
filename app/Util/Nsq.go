package Util

import (
	"api-skeleton/app/Global"
	"fmt"
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
