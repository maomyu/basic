package redispool

import (
	"fmt"
	"io"

	"github.com/garyburd/redigo/redis"
)

// 新创建一个连接会话
func NewSession() (*redisConnection, error, *Pool, io.Closer) {
	conn, err := p.Acquire()
	if err != nil {
		return nil, fmt.Errorf("[basic]：获取一个连接失败,[%w]", err), p, conn.(*redisConnection)
	}
	fa := conn.(*redisConnection)
	return fa, nil, p, conn.(*redisConnection)
}
func (s *redisConnection) Relase(p *Pool, c io.Closer) {
	p.Release(c)
}
func (f *redisConnection) GetID() int32 {
	return f.id
}

func (f *redisConnection) GetConn() redis.Conn {
	return f.conn
}
