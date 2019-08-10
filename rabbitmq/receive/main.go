/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-09 16:19:55
 * @LastEditTime: 2019-08-09 19:22:55
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/streadway/amqp"
	client "github.com/yuwe1/common/rabbitmq/client"
)

var forever = make(chan bool)

func recive(d amqp.Delivery) {
	fmt.Println(string(d.Body))
	d.Acknowledger.Ack(d.DeliveryTag, true)
}
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

	go testBroadQueueType().ConsumeFromQueue(recive)
	http.ListenAndServe("0.0.0.0:8200", nil)
}
