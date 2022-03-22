package Consumer

import (
	"api-skeleton/app/Consumer/KafkaConsumer"
	"api-skeleton/app/Consumer/NsqConsumer"
	"api-skeleton/bootstrap"
	"api-skeleton/config"
	"fmt"
)

var configs = config.InitConfig

var (
	nsqConsumer   *NsqConsumer.Consumer
	kafkaConsumer *KafkaConsumer.Consumer
	address       = fmt.Sprintf("%s:%s", configs.Nsq.Host, configs.Nsq.LookUpNode)
	kafkaAddr     = []string{
		configs.Kafka.Node1,
	}
)

type MsgType struct {
	Cate, Topic, Chanel string
}

//NewMsgType 消息类型构造函数
func NewMsgType(cate, topic, chanel string) *MsgType {
	return &MsgType{
		Cate:   cate,
		Topic:  topic,
		Chanel: chanel,
	}
}

//
func (m MsgType) InitMsgConsumer() {
	//初始化配置
	bootstrap.InitConfig()
	switch m.Cate {
	case "nsq":
		initNsqConsumer(m.Topic, m.Chanel)
	case "kafka":
		initKafkaConsumer(m.Topic)
	default:
		initNsqConsumer(m.Topic, m.Chanel)
	}
}

//initNsqConsumer 初始化nsq客户端配置
func initNsqConsumer(topic string, chanel string) {
	nsqConsumer = new(NsqConsumer.Consumer)
	nsqConsumer.Address = address
	nsqConsumer.Topic = topic
	nsqConsumer.Chanel = chanel
}

func initKafkaConsumer(topic string) {
	kafkaConsumer = KafkaConsumer.NewKafkaConsumer(topic, kafkaAddr, int32(2))
}
