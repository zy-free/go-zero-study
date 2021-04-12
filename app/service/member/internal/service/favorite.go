package service

import (
	"context"
	"go-zero-study/app/service/member/api"
	"go-zero-study/app/service/member/internal/model"
	"go-zero-study/core/logx"
)

type FavoriteLogic struct {
	ctx context.Context
	svc *Service
	logx.Logger
}

func NewFavoriteLogic(ctx context.Context, svc *Service) *FavoriteLogic {
	return &FavoriteLogic{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteLogic) AddFavorite(in *api.AddFavoriteReq) (*api.IDResp, error) {
	id, err := l.svc.FavoriteDao.Insert(model.Favorite{
		Mid:  in.Mid,
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}

	return &api.IDResp{Id: id}, nil
}
