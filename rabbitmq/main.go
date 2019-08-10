/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-09 16:11:05
 * @LastEditTime: 2019-08-09 19:22:14
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"encoding/json"
	"fmt"

	client "github.com/yuwe1/common/rabbitmq/client"
)

func testSimple() *client.MsgClient {
	simplequeue := client.NewSimpleQueue("user", "Login", true)
	body, _ := json.Marshal(simplequeue)
	fmt.Println(string(body))
	Simple := &client.MsgClient{
		Type: client.SimpleQueueType,
		Data: string(body),
	}
	body, _ = json.Marshal(Simple)
	fmt.Println(string(body))
	Simple.ConnectToRabbitmq("amqp://admin:admin@192.168.10.252:5672")
	fmt.Println(Simple)
	return Simple
}
func testBroadQueueType() *client.MsgClient {
	broadqueue := client.NewComplexQueue("broadqueue_exchange", "broadqueue_route", "", true)
	body, _ := json.Marshal(broadqueue)
	Simple := &client.MsgClient{
		Type: client.BroadQueueType,
		Data: string(body),
	}
	Simple.ConnectToRabbitmq("amqp://admin:admin@192.168.10.252:5672")
	fmt.Println(Simple)
	return Simple
}
func main() {
	msg := "hello"
	// testSimple().PublishToQueue([]byte(msg))
	testBroadQueueType().PublishToQueue([]byte(msg))
}
