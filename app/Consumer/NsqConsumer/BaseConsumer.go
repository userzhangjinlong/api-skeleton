package NsqConsumer

import (
	"api-skeleton/app/Global"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

type Consumer struct {
	Topic, Chanel, Address string
	cfg                    *nsq.Config
}

type MessageHandler struct {
	nsqConsumer *nsq.Consumer
}

//NewNsqConfig 初始化nsq配置
func (c *Consumer) NewNsqConfig() {
	c.cfg = nsq.NewConfig()
	c.cfg.AuthSecret = Global.Configs.Nsq.Password
	c.cfg.LookupdPollInterval = 3 * time.Second
	return
}

//HandleMessage 统一管理接收消息消费处理逻辑
func (m *MessageHandler) HandleMessage(msg *nsq.Message) error {
	//这里做消息消费处理逻辑
	fmt.Printf("接受到的消息体内容是:%s", string(msg.Body))
	return nil
}
