package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	MySQL   struct {
		DataSource string
	}
	Redis struct {
		Addr string
	}
	COS struct {
		Bucket    string
		SecretID  string
		SecretKey string
	}
}
