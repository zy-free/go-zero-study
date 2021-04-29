package config

import (
	"go-zero-study/rest"
	"go-zero-study/zrpc"
)

type Config struct {
	rest.RestConf
	Comment zrpc.RpcClientConf
}
