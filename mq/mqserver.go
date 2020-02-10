package mq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 定义rabbitmq的接口方法
type IMMessageClient interface {
	ConnectToRabbitmq(Connection string) *amqp.Connection
	PublishOnQueue(msg []byte, exchangename string, consumerName string, queueName string) error
	ConsumeFromQueue(queueName string, consumerName string, handlerfunc func(d amqp.Delivery)) error
	Close()
}
type messagingClient struct {
	Conn *amqp.Connection
}

func (m *messagingClient) ConnectToRabbitmq(Connection string) *amqp.Connection {
	var err error
	m.Conn, err = amqp.Dial(fmt.Sprintf("%s/", Connection))
	if err != nil {
		log.Fatal(err)
	}
	return m.Conn
}
func (m *messagingClient) PublishOnQueue(msg []byte, exchangename string, consumerName string, queueName string) error {
	if m.Conn != nil {
		ch, err := m.Conn.Channel()
		FailOnError(err, "Failed to open a channel")
		defer ch.Close()

		err = ch.ExchangeDeclare(
			exchangename, // name交换器的名字
			"direct",     // type交换器的类型
			true,         // durable是否持久化
			false,        // auto-deleted没有队列绑定时是否自动删除
			false,        // internal
			false,        // no-wait
			nil,          // arguments
		)
		q, err := ch.QueueDeclare(
			queueName, //name队列的名称
			true,      //durble是否持久化
			false,     //delete when unused是否自动删除
			false,     //exclusive是否设置排他，如果设置为true，则队列仅对首次声明他的连接可见，并在连接断开的时候自动删除
			false,     //no-wait是否阻塞
			nil,       //arguments
		)
		// 队列绑定
		ch.QueueBind(
			q.Name,
			consumerName,
			exchangename,
			false,
			nil,
		)
		err = ch.Publish(
			exchangename,
			consumerName,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         msg,
			})
		FailOnError(err, "fail to publish a message")
		log.Printf(" [x] Sent %s", msg)
	}
	return nil
}

var forever = make(chan bool)

func (m *messagingClient) ConsumeFromQueue(queueName string, consumerName string, handlerfunc func(d amqp.Delivery)) error {
	if m.Conn != nil {
		ch, err := m.Conn.Channel()
		FailOnError(err, "Failed to open a channel")
		defer ch.Close()
		q, err := ch.QueueDeclare(
			queueName, //name队列的名称
			true,      //durble是否持久化
			false,     //delete when unused是否自动删除
			false,     //exclusive是否设置排他，如果设置为true，则队列仅对首次声明他的连接可见，并在连接断开的时候自动删除
			false,     //no-wait是否阻塞
			nil,       //arguments
		)

		msgs, err := ch.Consume(
			q.Name,       // queue
			consumerName, // consumer
			false,        // auto-ack
			false,        // exclusive
			false,        // no-local
			false,        // no-wait
			nil,          // args
		)
		FailOnError(err, "Failed to  register a consumer")
		fmt.Println(msgs)

		fmt.Println("准备读取数据")
		go consumeLoop(msgs, handlerfunc)
		fmt.Println(<-forever)

	}
	fmt.Println("结束程序")
	return nil
}
func (m *messagingClient) Close() {
	if m.Conn != nil {
		m.Conn.Close()
	}
}

// var p = make(chan bool)
func consumeLoop(deliveries <-chan amqp.Delivery, handlerfunc func(d amqp.Delivery)) {

	for d := range deliveries {
		fmt.Println("有数据:", string(d.Body))
		handlerfunc(d)
	}
	forever <- true
	fmt.Println("chulaiel")
}
