package service

import (
	"context"
	"errors"
	"go-zero-study/app/service/comments/internal/dao/comment"
	"go-zero-study/core/logx"

	"go-zero-study/app/service/comments/api"
)

type CommentsLogic struct {
	ctx    context.Context
	svcCtx *Service
	logx.Logger
}

func NewCommentsLogic(ctx context.Context, svcCtx *Service) *CommentsLogic {
	return &CommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentsLogic) BathComments(in *api.BatchCommentsReq) (*api.BatchCommentsResp, error) {
	count, err := l.svcCtx.CommentDao.Count(int64(in.RoomId))
	if err != nil {
		return nil, err
	}

	page := in.Page
	size := in.Size
	if count <= (page-1) * size {
		return nil, errors.New("not fund")
	}

	start := (page-1) * size
	comments, err := l.svcCtx.CommentDao.BatchFindByRoomId(int64(in.RoomId), int64(start), int64(size))
	if err != nil {
		return nil, err
	}

	var cms []*api.Comment
	for _, comment := range comments {
		cms = append(cms, &api.Comment{
			Id: comment.Id,
			UserId: comment.UserId,
			Content: comment.Content,
			Nickname: comment.Nickname,
			Avatar: comment.Avatar,
			RoomId: comment.RoomId,
		})
	}

	resp := &api.BatchCommentsResp{
		TotalPage: (count-1)/size + 1,
		TotalSize: count,
		Comments: cms,
	}



	return resp, nil
}

func (l *CommentsLogic) CreateComment(in *api.CreateCommentReq) (*api.CreateCommentResp, error) {
	result, err := l.svcCtx.CommentDao.Insert(comment.Comments{
		UserId: int64(in.UserId),
		Avatar: in.Avatar,
		RoomId: int64(in.RoomId),
		Nickname: in.Nickname,
		Content: in.Content,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &api.CreateCommentResp{
		Id: id,
	}, nil
}
