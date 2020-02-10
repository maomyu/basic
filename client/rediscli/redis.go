package rediscli

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/basic/config"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		fmt.Println("已经初始化redis客户端")
		return
	}
	redisConfig := config.GetRedisConfig()

	// 加载配置文件成redis客户端
	if redisConfig != nil && redisConfig.GetEnabled() {
		if redisConfig.GetSentinelConfig() != nil && redisConfig.GetSentinelConfig().GetEnabled() {
			initSentinel(redisConfig)
		} else {
			initSingle(redisConfig)
		}
	}
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Log("初始化Redis，检测连接Ping.")
	log.Log("初始化Redis，检测连接Ping..")
	log.Logf("初始化Redis，检测连接Ping... %s", pong)
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSentinel(redisConfig config.RedisConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})

}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
