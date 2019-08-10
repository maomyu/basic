/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-09 15:17:12
 * @LastEditTime: 2019-08-09 19:19:15
 * @LastEditors: Please set LastEditors
 */
package client

import (
	"log"

	"github.com/streadway/amqp"
)

// 定义rabbitMQ的接口方法
type IMMessageClient interface {
	// 连接RabbitMQ，并获取连接
	ConnectToRabbitmq(Connection string)
	// 发送消息
	PublishToQueue(msg []byte) error
	// 消费消息
	ConsumeFromQueue(handlerfunc func(d amqp.Delivery)) error
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

const (
	SimpleQueueType = "SimpleQueue"
	BroadQueueType  = "BroadQueue"
	DirectQueueType = "DirectQueue"
	TopicQueueType  = "TopicQueue"
)

type MsgClient struct {
	Conn *amqp.Connection
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //队列数据
}
type SimpleQueue struct {
	Rout_key      string `json:"rout_key"`      //路由
	Queue         string `json:"queue"`         //队列的名字
	Is_persistent bool   `json:"is_persistent"` //队列是否持久化
}

func NewSimpleQueue(rout_key string, queue string, is_persistent bool) *SimpleQueue {
	return &SimpleQueue{
		Rout_key:      rout_key,
		Queue:         queue,
		Is_persistent: is_persistent,
	}
}

type ComplexQueue struct {
	ExchangeName  string `json:"exchangeName"`
	Rout_key      string `json:"rout_key"`      //路由
	Queue         string `json:"queue"`         //队列的名字
	Is_persistent bool   `json:"is_persistent"` //队列是否持久化
}

func NewComplexQueue(exchangeName string, rout_key string, queue string, is_persistent bool) *ComplexQueue {
	return &ComplexQueue{
		ExchangeName:  exchangeName,
		Rout_key:      rout_key,
		Queue:         queue,
		Is_persistent: is_persistent,
	}
}

type TopicQueue struct {
	ExchangeName  string `json:"exchangeName"`
	Rout_key      string `json:"rout_key"`      //路由
	Queue         string `json:"queue"`         //队列的名字
	Is_persistent bool   `json:"is_persistent"` //队列是否持久化
	Bind_key      string `json:"bind_key"`      //绑定的路由
}

func NewTopicQueue(exchangeName string, rout_key string, bind_key, string, queue string, is_persistent bool) *TopicQueue {
	return &TopicQueue{
		ExchangeName:  exchangeName,
		Rout_key:      rout_key,
		Bind_key:      bind_key,
		Queue:         queue,
		Is_persistent: is_persistent,
	}
}
