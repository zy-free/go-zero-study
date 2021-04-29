
package api

import (
	"context"
	"go-zero-study/zrpc"
)

type (
	CommentsRPC interface {
		BathComments(ctx context.Context, in *BatchCommentsReq) (*BatchCommentsResp, error)
		CreateComment(ctx context.Context, in *CreateCommentReq) (*CreateCommentResp, error)
	}

	defaultCommentsRPC struct {
		cli zrpc.Client
	}
)

func NewCommentsRPC(cli zrpc.Client) CommentsRPC {
	return &defaultCommentsRPC{
		cli: cli,
	}
}

func (m *defaultCommentsRPC) BathComments(ctx context.Context, in *BatchCommentsReq) (*BatchCommentsResp, error) {
	client := NewCommentsRPCClient(m.cli.Conn())
	return client.BathComments(ctx, in)
}

func (m *defaultCommentsRPC) CreateComment(ctx context.Context, in *CreateCommentReq) (*CreateCommentResp, error) {
	client := NewCommentsRPCClient(m.cli.Conn())
	return client.CreateComment(ctx, in)
}
