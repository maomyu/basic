package config

import "strings"

type RedisConfig interface {
	GetEnabled() bool
	GetConn() string
	GetPassword() string
	GetDBNum() int
	GetSentinelConfig() RedisSentinelConfig
}

type RedisSentinelConfig interface {
	GetEnabled() bool
	GetMaster() string
	GetNodes() []string
}
type defaultRedisConfig struct {
	Enabled  bool          `json:"enabled",yaml:"enabled"`
	Conn     string        `json:"conn",yaml:"conn"`
	Password string        `json:"password",yaml:"password"`
	DBNum    int           `json:"dbnum",yaml:"dbnum"`
	Timeout  int           `json:"timeout",yaml:"timeout"`
	sentinel redisSentinel `json:"sentinel",yaml:"sentinel"`
}
type redisSentinel struct {
	Enabled bool   `json:"enabled",yaml:"enabled"`
	Master  string `json:"master",yaml:"master"`
	Nodes   string `json:"nodes",yaml:"nodes"`
	nodes   []string
}

// GetEnabled redis 配置是否激活
func (r defaultRedisConfig) GetEnabled() bool {
	return r.Enabled
}

// GetConn redis 地址
func (r defaultRedisConfig) GetConn() string {
	return r.Conn
}

// GetPassword redis 密码
func (r defaultRedisConfig) GetPassword() string {
	return r.Password
}

// GetDBNum redis 数据库分区序号
func (r defaultRedisConfig) GetDBNum() int {
	return r.DBNum
}

// GetDBNum redis 数据库分区序号
func (r defaultRedisConfig) GetSentinelConfig() RedisSentinelConfig {
	return r.sentinel
}

// GetEnabled redis 哨兵配置是否激活
func (s redisSentinel) GetEnabled() bool {
	return s.Enabled
}

// GetMaster redis 主节点名
func (s redisSentinel) GetMaster() string {
	return s.Master
}

// GetNodes redis 哨兵节点列表
func (s redisSentinel) GetNodes() []string {
	if len(s.Nodes) != 0 {
		for _, v := range strings.Split(s.Nodes, ",") {
			v = strings.TrimSpace(v)
			s.nodes = append(s.nodes, v)
		}
	}

	return s.nodes
}
