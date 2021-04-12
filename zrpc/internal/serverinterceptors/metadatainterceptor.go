package serverinterceptors

import (
	"context"
	nmd "go-zero-study/core/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

func UnaryMetaDataInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		var (
			caller     string
			color      string
			remote     string
			remotePort string
			mid        int64
		)
		if gmd, ok := metadata.FromIncomingContext(ctx); ok {
			if strs, ok := gmd[nmd.Color]; ok {
				color = strs[0]
			}
			if strs, ok := gmd[nmd.RemoteIP]; ok {
				remote = strs[0]
			}
			if callers, ok := gmd[nmd.Caller]; ok {
				caller = callers[0]
			}
			if remotePorts, ok := gmd[nmd.RemotePort]; ok {
				remotePort = remotePorts[0]
			}
			if mids, ok := gmd[nmd.Mid]; ok {
				midInt, _ := strconv.Atoi(mids[0])
				mid = int64(midInt)
			}
		}

		ctx = nmd.NewContext(ctx, nmd.MD{
			nmd.Color:      color,
			nmd.RemoteIP:   remote,
			nmd.Caller:     caller,
			nmd.RemotePort: remotePort,
			nmd.Mid:        mid,
		})
		return handler(ctx, req)
	}
}
