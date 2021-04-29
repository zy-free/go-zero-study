package service

import (
	"go-zero-study/app/api/comments/config"
	"go-zero-study/app/service/comments/api"
	"go-zero-study/zrpc"
)

type ServiceContext struct {
	Config config.Config
	CommentRPC api.CommentsRPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		CommentRPC: api.NewCommentsRPC(zrpc.MustNewClient(c.Comment)),
	}
}
