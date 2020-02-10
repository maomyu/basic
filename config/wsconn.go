package config

type WsConf interface {
	GetWsListenAddr() string
	GetRpcListenAddr() string
	GetLocalAddr() string
	GetLogicRPCAddrs() string
}

type defaultWsConfig struct {
	WsListenAddr  string `yml:"wsListenAddr"`
	RpcListenAddr string `yml:"rpcListenAddr"`
	LocalAddr     string `yml:"localAddr"`
	LogicRPCAddrs string `yml:"logicRPCAddrs"`
}

func (l defaultWsConfig) GetWsListenAddr() string {
	return l.WsListenAddr
}
func (l defaultWsConfig) GetRpcListenAddr() string {
	return l.RpcListenAddr
}
func (l defaultWsConfig) GetLocalAddr() string {
	return l.LocalAddr
}
func (l defaultWsConfig) GetLogicRPCAddrs() string {
	return l.LogicRPCAddrs
}
