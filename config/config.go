package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
)

var (
	err error
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	redisConfig             defaultRedisConfig
	profiles                defaultProfiles
	rabbitMQConfig          defaultRabbitMQConfig
	mysqlConfig             defaultMysqlConfig
	logicConfig             defaultLogicConfig
	wsconnfig               defaultWsConfig
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

func Init() {
	m.Lock()
	defer m.Unlock()
	if inited {
		log.Logf("配置文件已经被初始化过了")
	}
	//获取当前绝对路径
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))
	// 获得conf的路径
	pt := filepath.Join(appPath, "./../srv_conf")
	//设置工作目录
	fmt.Println(pt)
	// os.Chdir(filepath.Join(pt, "."))
	// 加载我们的配置文件
	fmt.Println(pt)
	if err = config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}

	// 获取值并进行赋值
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}
	log.Logf("[Init] 加载配置文件：path: %s, %+v\n", pt+sp+"application.yml", profiles)
	// 根据profiles；来判断有几个配置文件需要进行加载
	if len(profiles.GetInclude()) > 0 {
		// 找到所有的配置文件，使用字符串切割
		include := strings.Split(profiles.GetInclude(), ",")
		// 根据include的容量来定义一个切片
		sources := make([]source.Source, len(include))
		// 和之前的加载模式一样，这次需要用到循环
		for i := 0; i < len(include); i++ {
			// 获得准备加载文件的路径
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			sources[i] = file.NewSource(file.WithPath(filePath))
		}
		// 循环结束后开始加载文件
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}
	// 完成配置文件的赋值操作
	config.Get(defaultRootPath, "redis").Scan(&redisConfig)
	config.Get(defaultRootPath, "rabbitmq").Scan(&rabbitMQConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)
	config.Get(defaultRootPath, "logic").Scan(&logicConfig)
	config.Get(defaultRootPath, "wsconn").Scan(&wsconnfig)
	inited = true
}

// GetRedisConfig 获取Redis配置
func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}
func GetrabbitMQConfig() (ret RabbitMQConfig) {

	return rabbitMQConfig
}

//获取我们的配置
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// 获取logic的配置
func GetLogicConfig() (ret LogicConf) {
	return logicConfig
}

func GetWsconfConfig() (ret WsConf) {
	return wsconnfig
}
