package rpc

import (
	"context"
	"go-zero-study/app/service/member/api"
	"go-zero-study/app/service/member/internal/service"
)

func (s *Server) AddFavorite(ctx context.Context, in *api.AddFavoriteReq) (*api.IDResp, error) {
	l := service.NewFavoriteLogic(ctx,s.svc)
	return l.AddFavorite(in)
}