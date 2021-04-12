package model

import (
	"go-zero-study/core/stores/cache"
	"go-zero-study/core/stores/sqlc"
	"go-zero-study/core/stores/sqlx"
)


type Dao struct {
	sqlc.CachedConn
	table string
}

func New(conn sqlx.SqlConn, c cache.CacheConf, table string) *Dao {

	return &Dao{
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}
