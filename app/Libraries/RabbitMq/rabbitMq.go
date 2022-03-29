package RabbitMq

import (
	"api-skeleton/app/Global"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	coon    *amqp.Connection
	channel *amqp.Channel
	//队列名称 交换机 key 连接信息
	QueueName, Exchange, Key, Mqurl string
}

//NewRabbitMQ 构造rabbitMq实例
func NewRabbitMQ(queue, exchange, key string) *RabbitMQ {
	// MQURL 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost(访问用户)
	mqurl := fmt.Sprintf("amqp://%s:%s@%s",
		Global.Configs.Rabbitmq.Username,
		Global.Configs.Rabbitmq.Password,
		Global.Configs.Rabbitmq.Node1,
		//Global.Configs.Rabbitmq.Vhost,
	)
	rabbitMq := &RabbitMQ{
		QueueName: queue,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     mqurl,
	}

	var err error
	rabbitMq.coon, err = amqp.Dial(rabbitMq.Mqurl)
	rabbitMq.recoverErr(err, "connect")

	rabbitMq.channel, err = rabbitMq.coon.Channel()
	rabbitMq.recoverErr(err, "channel")

	return rabbitMq
}

//ProducerMsg 生产消息推送
func (r *RabbitMQ) ProducerMsg(msg string) {
	defer r.Destory()
	q, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	r.recoverErr(err, "failed to declare a queue")

	err = r.channel.Publish(r.Exchange, q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
	r.recoverErr(err, "failed to producer msg")
}

//Destory 销毁channel和coon 链接
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.coon.Close()
}

func (r *RabbitMQ) recoverErr(err error, message string) {
	if err != nil {
		logrus.Errorf("rabbitMq %s err:%s", message, err)
		panic(fmt.Sprintf("rabbitMq %s err", message))
	}
}
