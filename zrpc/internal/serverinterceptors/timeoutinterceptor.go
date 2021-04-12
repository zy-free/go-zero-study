package serverinterceptors

import (
	"context"
	"github.com/pkg/errors"
	"time"

	"go-zero-study/core/contextx"
	"google.golang.org/grpc"
)

func UnaryTimeoutInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, cancel := contextx.ShrinkDeadline(ctx, timeout)
		defer cancel()
		resp, err =  handler(ctx, req)
		err = errors.Cause(err)
		return resp, err

	}
}
