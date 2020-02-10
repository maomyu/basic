package dbpool

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yuwe1/basic/config"
	"github.com/yuwe1/basic/logger"
)

var (
	inited  bool
	mysqlDB *sql.DB
	p       *Pool
	m       sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()
	var err error
	if inited {
		err = fmt.Errorf("[init]db已经被初始化")
		logger.Sugar.Errorf(err.Error())
		return
	}
	fmt.Println(config.GetMysqlConfig())
	if config.GetMysqlConfig().GetEnabled() {
		initPool()
	}
	inited = true
}
func initPool() {
	p, _ = New(createConnection, config.GetMysqlConfig().GetMaxOpenConnection())
}

//数据库连接
type dbConnection struct {
	Db *sql.DB
	ID int32 //连接的标志
}

//实现io.Closer接口
func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCounter int32

//生成数据库连接的方法，以供资源池使用
func createConnection() (io.Closer, error) {
	//并发安全，给数据库连接生成唯一标志
	id := atomic.AddInt32(&idCounter, 1)
	fmt.Println(config.GetMysqlConfig())
	mysqlDB, _ := sql.Open("mysql", config.GetMysqlConfig().GetURL())
	return &dbConnection{
		ID: id,
		Db: mysqlDB,
	}, nil
}
