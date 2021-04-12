package service

import (
	"go-zero-study/app/service/member/config"
	"go-zero-study/app/service/member/internal/dao/favorite"
	memberDao "go-zero-study/app/service/member/internal/dao/member"

	"go-zero-study/core/stores/sqlx"
)

type Service struct {
	c           config.Config
	MemberDao   *memberDao.Dao
	FavoriteDao *favorite.Dao
	// more
}

func New(c config.Config) *Service {
	sqlConn := sqlx.NewMysql(c.DataSource)
	return &Service{
		c:           c,
		MemberDao:   memberDao.New(sqlConn, c.Cache, c.MemberTable),
		FavoriteDao: favorite.New(sqlConn,c.FavoriteTable),
		// more
	}
}
