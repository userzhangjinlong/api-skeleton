package NsqConsumer

import (
	"api-skeleton/app/Util"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (c *Consumer) CreateRankingListConsumer() {
	c.NewNsqConfig()
	consumer, err := Util.CreateNsqConsumer(c.Topic, c.Chanel, c.cfg)
	fmt.Printf("topic:%s, chanel:%s", c.Topic, c.Chanel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Err":    err.Error(),
			"Topic":  c.Topic,
			"Chanel": c.Chanel,
		}).Errorf("CreateRankingListConsumer消费者声明异常:%s", err.Error())
	}

	go func() {
		handler := MessageHandler{nsqConsumer: consumer}
		consumer.AddHandler(&handler)
		err = consumer.ConnectToNSQLookupd(c.Address)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Err": err.Error(),
			}).Errorf("CreateRankingListConsumer消费者链接异常:%s", err.Error())
		}
	}()
	select {}

}
