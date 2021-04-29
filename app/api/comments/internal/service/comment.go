package service

import (
	"context"
	"go-zero-study/app/api/comments/internal/model"
	"go-zero-study/app/service/comments/api"
	"go-zero-study/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *ServiceContext) CommentLogic {
	return CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) CreateComment(req model.CreateCommentReq) (*model.CreateCommentResp, error) {
	content := req.Content

	r := &api.CreateCommentReq{
		UserId: 1,
		RoomId: 1,
		Nickname: "张三",
		Avatar: "",
		Content:  content,
	}

	_, err := l.svcCtx.CommentRPC.CreateComment(l.ctx, r)
	if err != nil {
		return nil, err
	}

	return &model.CreateCommentResp{}, nil
}

func (l *CommentLogic) BatchQueryComments(req model.BatchQueryCommentsReq) (*model.BathQueryCommentsResp, error) {
	r := &api.BatchCommentsReq{
		RoomId: int64(req.RoomId),
		Page: int64(req.Page),
		Size: int64(req.Size),
	}

	p, err := l.svcCtx.CommentRPC.BathComments(l.ctx, r)
	if err != nil {
		return nil, err
	}

	var resp model.BathQueryCommentsResp
	resp.TotalPage = int(p.TotalPage)
	resp.TotalSize = int(p.TotalSize)

	for _, comment := range p.Comments {
		resp.Comments = append(resp.Comments, &model.Comment{
			Id: int(comment.Id),
			UserId: int(comment.UserId),
			Nickname: comment.Nickname,
			Avatar: comment.Avatar,
			RoomId: int(comment.RoomId),
			Content: comment.Content,
		})
	}

	return &resp, nil
}