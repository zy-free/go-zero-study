package service

import (
	"go-zero-study/app/service/comments/config"
	"go-zero-study/app/service/comments/internal/dao/comment"
	"go-zero-study/core/stores/sqlx"
)

type Service struct {
	Config config.Config
	CommentDao  comment.CommentsModel
}

func NewServiceContext(c config.Config) *Service {
	sqlConn := sqlx.NewMysql(c.DataSource)
	return &Service{
		Config: c,
		CommentDao: comment.NewCommentsModel(sqlConn, c.Cache, c.CommentTable),
	}
}
