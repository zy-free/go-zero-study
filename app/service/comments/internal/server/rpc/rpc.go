package rpc

import (
	"context"
	"go-zero-study/app/service/comments/internal/service"

	"go-zero-study/app/service/comments/api"
)

type CommentsRPCServer struct {
	svcCtx *service.Service
}

func NewCommentsRPCServer(svcCtx *service.Service) *CommentsRPCServer {
	return &CommentsRPCServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentsRPCServer) BathComments(ctx context.Context, in *api.BatchCommentsReq) (*api.BatchCommentsResp, error) {
	l := service.NewCommentsLogic(ctx, s.svcCtx)
	return l.BathComments(in)
}

func (s *CommentsRPCServer) CreateComment(ctx context.Context, in *api.CreateCommentReq) (*api.CreateCommentResp, error) {
	l := service.NewCommentsLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}
