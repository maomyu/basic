package basic

import (
	"github.com/yuwe1/basic/client/dbpool"
	"github.com/yuwe1/basic/client/rediscli/redispool"
	"github.com/yuwe1/basic/config"
	"github.com/yuwe1/basic/mq"
)

func Init() {
	config.Init()
	// rediscli.Init()
	redispool.Init()
	mq.Init()
	dbpool.Init()
}
