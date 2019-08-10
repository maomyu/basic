/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-09 15:19:31
 * @LastEditTime: 2019-08-09 20:10:58
 * @LastEditors: Please set LastEditors
 */
package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// var p = make(chan bool)
func consumeLoop(deliveries <-chan amqp.Delivery, handlerfunc func(d amqp.Delivery)) {

	for d := range deliveries {
		fmt.Println("有数据:", string(d.Body))
		handlerfunc(d)
	}
}
func (m *MsgClient) ConnectToRabbitmq(Connection string) {

	var err error
	m.Conn, err = amqp.Dial(fmt.Sprintf("%s/", Connection))
	if err != nil {
		log.Fatal(err)
	}
}
func (m *MsgClient) PublishToQueue(msg []byte) error {
	if m.Conn != nil {
		ch, err := m.Conn.Channel()
		FailOnError(err, "Failed to open a channel")
		defer ch.Close()
		if m.Type == SimpleQueueType {
			var s SimpleQueue
			json.Unmarshal([]byte(m.Data), &s)
			q, err := ch.QueueDeclare(
				s.Queue,         //name队列的名称
				s.Is_persistent, //durble是否持久化
				false,           //delete when unused是否自动删除
				false,           //exclusive是否设置排他，如果设置为true，则队列仅对首次声明他的连接可见，并在连接断开的时候自动删除
				false,           //no-wait是否阻塞
				nil,             //arguments
			)
			FailOnError(err, "队列申请失败")
			err = ch.Publish(
				"",
				q.Name, // 路由，即队列的名字
				false,  //mandatory
				false,  //immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent, //消息的持久化
					ContentType:  "text/plain",
					Body:         msg,
				},
			)
			FailOnError(err, "发送消息失败")
		} else if m.Type == DirectQueueType {
			var s ComplexQueue
			json.Unmarshal([]byte(m.Data), &s)
			err = ch.ExchangeDeclare(
				s.ExchangeName, //交换器的名字
				"direct",       //交换器的类型、这里为广播类型
				true,           //是否持久
				false,          //无用的时候是否自动删除
				false,          //true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
				false,          //no-wait
				nil,            //arguments
			)
			err = ch.Publish(
				s.ExchangeName, //发送到交换机的名字
				s.Rout_key,     // 路由,需要消费者绑定
				false,          //mandatory
				false,          //immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent, //消息的持久化
					ContentType:  "text/plain",
					Body:         msg,
				},
			)

		} else if m.Type == BroadQueueType {
			var s ComplexQueue
			json.Unmarshal([]byte(m.Data), &s)
			err = ch.ExchangeDeclare(
				s.ExchangeName, //交换器的名字
				"fanout",       //交换器的类型、这里为广播类型
				true,           //是否持久
				false,          //无用的时候是否自动删除
				false,          //true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
				false,          //no-wait
				nil,            //arguments
			)
			err = ch.Publish(
				s.ExchangeName, //发送到交换机的名字
				"",             // 路由,对于广播模式，路由可以省略
				false,          //mandatory
				false,          //immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent, //消息的持久化
					ContentType:  "text/plain",
					Body:         msg,
				},
			)
		} else {
			var s TopicQueue
			json.Unmarshal([]byte(m.Data), &s)
			err = ch.ExchangeDeclare(
				s.ExchangeName, //交换器的名字
				"topic",        //交换器的类型、这里为广播类型
				true,           //是否持久
				false,          //无用的时候是否自动删除
				false,          //true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
				false,          //no-wait
				nil,            //arguments
			)
			err = ch.Publish(
				s.ExchangeName, //发送到交换机的名字
				s.Rout_key,     // 路由,需要消费者绑定
				false,          //mandatory
				false,          //immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent, //消息的持久化
					ContentType:  "text/plain",
					Body:         msg,
				},
			)
		}
	}
	return nil
}
func (m *MsgClient) ConsumeFromQueue(handlerfunc func(d amqp.Delivery)) error {
	if m.Conn != nil {
		ch, err := m.Conn.Channel()
		FailOnError(err, "Failed to open a channel")
		defer ch.Close()
		if m.Type == SimpleQueueType {
			var s SimpleQueue
			json.Unmarshal([]byte(m.Data), &s)
			q, err := ch.QueueDeclare(
				s.Queue,         //name队列的名称
				s.Is_persistent, //durble是否持久化
				false,           //delete when unused是否自动删除
				false,           //exclusive是否设置排他，如果设置为true，则队列仅对首次声明他的连接可见，并在连接断开的时候自动删除
				false,           //no-wait是否阻塞
				nil,             //arguments
			)
			FailOnError(err, "队列申请失败")
			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,  // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			consumeLoop(msgs, handlerfunc)

		} else if m.Type == DirectQueueType {
			var s ComplexQueue
			json.Unmarshal([]byte(m.Data), &s)
			err = ch.ExchangeDeclare(
				s.ExchangeName, //交换器的名字
				"direct",       //交换器的类型、这里为广播类型
				true,           //是否持久
				false,          //无用的时候是否自动删除
				false,          //true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
				false,          //no-wait
				nil,            //arguments
			)
			// 申请一个队列
			q, _ := ch.QueueDeclare(
				s.Queue, //name
				true,    //durable
				false,   //delete when usused
				true,    // exclusive
				false,   //no-wait
				nil,     // arguments
			)
			// 队列绑定
			err = ch.QueueBind(
				q.Name,         //队列的名字
				s.Rout_key,     //routing key
				s.ExchangeName, //所绑定的交换器
				false,
				nil,
			)
			msgs, _ := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,  // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			consumeLoop(msgs, handlerfunc)

		} else if m.Type == BroadQueueType {
			var s ComplexQueue
			json.Unmarshal([]byte(m.Data), &s)
			err = ch.ExchangeDeclare(
				s.ExchangeName, //交换器的名字
				"fanout",       //交换器的类型、这里为广播类型
				true,           //是否持久
				false,          //无用的时候是否自动删除
				false,          //true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
				false,          //no-wait
				nil,            //arguments
			)
			// 申请一个队列
			q, _ := ch.QueueDeclare(
				s.Queue, //name
				true,    //durable
				false,   //delete when usused
				true,    // exclusive
				false,   //no-wait
				nil,     // arguments
			)
			// 队列绑定
			err = ch.QueueBind(
				q.Name,         //队列的名字
				"",             //routing key
				s.ExchangeName, //所绑定的交换器
				false,
				nil,
			)
			msgs, _ := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,  // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			consumeLoop(msgs, handlerfunc)
		} else {
			var s TopicQueue
			json.Unmarshal([]byte(m.Data), &s)
			err = ch.ExchangeDeclare(
				s.ExchangeName, //交换器的名字
				"topic",        //交换器的类型、这里为广播类型
				true,           //是否持久
				false,          //无用的时候是否自动删除
				false,          //true表示是。客户端无法直接发送msg到内部交换器，只有交换器可以发送msg到内部交换器。
				false,          //no-wait
				nil,            //arguments
			)
			// 申请一个队列
			q, _ := ch.QueueDeclare(
				s.Queue, //name
				true,    //durable
				false,   //delete when usused
				true,    // exclusive
				false,   //no-wait
				nil,     // arguments
			)
			// 队列绑定
			err = ch.QueueBind(
				q.Name,         //队列的名字
				s.Bind_key,     //routing key
				s.ExchangeName, //所绑定的交换器
				false,
				nil,
			)
			msgs, _ := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,  // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			consumeLoop(msgs, handlerfunc)
		}
	}
	return nil
}
