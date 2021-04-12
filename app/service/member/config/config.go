package config

import (
	"go-zero-study/core/stores/cache"
	"go-zero-study/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	MemberTable      string
	FavoriteTable      string
	Cache      cache.CacheConf
}
