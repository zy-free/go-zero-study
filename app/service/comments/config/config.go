package config

import (
	"go-zero-study/core/stores/cache"
	"go-zero-study/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Cache cache.ClusterConf
	DataSource string
	CommentTable string
}
