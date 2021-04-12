package clientinterceptors

import (
	"context"
	"go-zero-study/core/ecode"
	"google.golang.org/grpc"
	gstatus "google.golang.org/grpc/status"
)

func GrpcErrorInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	if err := invoker(ctx, method, req, reply, cc, opts...); err != nil {
		gst, _ := gstatus.FromError(err)
		err = ecode.String(gst.Message())
		return err
	}
	return nil
}
