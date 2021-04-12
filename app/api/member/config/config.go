package config

import (
	"go-zero-study/rest"
	"go-zero-study/zrpc"
)

type Config struct {
	rest.RestConf
	Member zrpc.RpcClientConf
}
