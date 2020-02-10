package mq

import (
	"fmt"
	"log"
	"sync"

	"github.com/yuwe1/basic/config"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

//mq服务定义的包
type MsgInterface interface{}

var (
	Client messagingClient
	m      sync.RWMutex
	inited bool
)

type Message struct {
}

func Init() {
	m.Lock()
	defer m.Unlock()
	if inited {
		fmt.Println("已经初始化了rabbitmq")
		return
	}
	rabbitmqconfig := config.GetrabbitMQConfig()
	// fmt.Println(rabbitmqconfig)
	if rabbitmqconfig != nil {
		InitRabbitMQ(rabbitmqconfig)
	}
}
func InitRabbitMQ(rabbitmqconfig config.RabbitMQConfig) {
	str := "amqp://" + rabbitmqconfig.GetUser() + ":" + rabbitmqconfig.GetPassword() + "@" + rabbitmqconfig.GetURL()
	Client.Conn = Client.ConnectToRabbitmq(str)
}
func GetRabbitMQ() messagingClient {
	return Client
}
