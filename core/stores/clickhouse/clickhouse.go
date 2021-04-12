package clickhouse

import (
	_ "github.com/ClickHouse/clickhouse-go"
	"go-zero-study/core/stores/sqlx"
)

const clickHouseDriverName = "clickhouse"

func New(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(clickHouseDriverName, datasource, opts...)
}
