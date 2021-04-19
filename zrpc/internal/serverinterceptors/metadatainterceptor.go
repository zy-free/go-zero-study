package serverinterceptors

import (
	"context"
	nmd "go-zero-study/core/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryMetaDataInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		cmd := nmd.MD{}
		if gmd, ok := metadata.FromIncomingContext(ctx); ok {
			for key, vals := range gmd {
				if nmd.IsIncomingKey(key) {
					cmd[key] = vals[0]
				}
			}
		}

		ctx = nmd.NewContext(ctx, cmd)
		return handler(ctx, req)
	}
}
