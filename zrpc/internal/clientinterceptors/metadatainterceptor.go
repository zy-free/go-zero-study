package clientinterceptors

import (
	"context"
	"fmt"
	nmd "go-zero-study/core/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func MetaDataInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	gmd  := metadata.MD{}

	nmd.Range(ctx,
		func(key string, value interface{}) {
			if valstr, ok := value.(string); ok {
				gmd[key] = []string{valstr}
			}
		},
		nmd.IsOutgoingKey)
	// merge with old matadata if exists
	if oldmd, ok := metadata.FromOutgoingContext(ctx); ok {
		gmd = metadata.Join(gmd, oldmd)
	}
	fmt.Println(gmd)
	ctx = metadata.NewOutgoingContext(ctx, gmd)

	return invoker(ctx, method, req, reply, cc, opts...)

}