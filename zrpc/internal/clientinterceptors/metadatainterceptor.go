package clientinterceptors

import (
	"context"
	nmd "go-zero-study/core/metadata"
	"google.golang.org/grpc/metadata"
	"strconv"

	"google.golang.org/grpc"
)

func MetaDataInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var(
		cmd    nmd.MD
		gmd    metadata.MD
		ok     bool
	)
	gmd = metadata.MD{}

	// meta color
	if cmd, ok = nmd.FromContext(ctx); ok {
		var color, ip, port string
		var mid int64
		if color, ok = cmd[nmd.Color].(string); ok {
			gmd[nmd.Color] = []string{color}
		}
		if ip, ok = cmd[nmd.RemoteIP].(string); ok {
			gmd[nmd.RemoteIP] = []string{ip}
		}
		if port, ok = cmd[nmd.RemotePort].(string); ok {
			gmd[nmd.RemotePort] = []string{port}
		}
		if mid, ok = cmd[nmd.Mid].(int64); ok {
			gmd[nmd.Mid] = []string{strconv.Itoa(int(mid))}
		}
	}
	// merge with old matadata if exists
	if oldmd, ok := metadata.FromOutgoingContext(ctx); ok {
		gmd = metadata.Join(gmd, oldmd)
	}
	ctx = metadata.NewOutgoingContext(ctx, gmd)

	return invoker(ctx, method, req, reply, cc, opts...)

}