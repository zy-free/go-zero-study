package comment

import (
	"go-zero-study/core/stores/cache"
	"go-zero-study/core/stores/sqlc"
	"go-zero-study/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

type defaultCommentsModel struct {
	sqlc.CachedConn
	table string
}

func NewCommentsModel(conn sqlx.SqlConn, c cache.CacheConf, table string) CommentsModel {
	return &defaultCommentsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}
