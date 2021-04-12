package favorite

import "go-zero-study/core/stores/sqlx"

type Dao struct {
	conn  sqlx.SqlConn
	table string
}

func New(conn sqlx.SqlConn,table string) *Dao {
	return &Dao{
		conn:  conn,
		table:      table,
	}
}
