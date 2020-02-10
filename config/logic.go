package config

type LogicConf interface {
	GetRpcIntListenAddr() string
	GetClientRpcExtListenAddr() string
	GetServerRpcExtListenAddr() string
	GetConnRpcAddrs() string
}

type defaultLogicConfig struct {
	RpcIntListenAddr       string `yml:"rpcIntListenAddr"`
	ClientRpcExtListenAddr string `yml:"clientRpcExtListenAddr"`
	ServerRpcExtListenAddr string `yml:"serverRpcExtListenAddr"`
	ConnRpcAddrs           string `yml:"connRpcAddrs"`
}

func (l defaultLogicConfig) GetRpcIntListenAddr() string {
	return l.RpcIntListenAddr
}

func (l defaultLogicConfig) GetClientRpcExtListenAddr() string {
	return l.ClientRpcExtListenAddr
}

func (l defaultLogicConfig) GetServerRpcExtListenAddr() string {
	return l.ServerRpcExtListenAddr
}

func (l defaultLogicConfig) GetConnRpcAddrs() string {
	return l.ConnRpcAddrs
}
