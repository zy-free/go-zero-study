package service

import (
	"go-zero-study/app/api/member/config"
	memberProto "go-zero-study/app/service/member/api"
	"go-zero-study/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	MemberCli memberProto.MemberCli
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		MemberCli: memberProto.New(zrpc.MustNewClient(c.Member)),
	}
}
