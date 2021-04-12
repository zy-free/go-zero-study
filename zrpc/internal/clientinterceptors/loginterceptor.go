package clientinterceptors

import (
	"context"
	"fmt"
	"go-zero-study/core/ecode"
	"google.golang.org/grpc/peer"
	"path"
	"time"

	"go-zero-study/core/logx"
	"go-zero-study/core/timex"
	"google.golang.org/grpc"
)

const slowThreshold = time.Millisecond * 500

func LogInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	serverName := path.Join(cc.Target(), method)

	start := timex.Now()
	var peerInfo peer.Peer
	opts = append(opts, grpc.Peer(&peerInfo))

	// invoker requests
	err := invoker(ctx, method, req, reply, cc, opts...)

	// after request
	code := ecode.Cause(err).Code()
	elapsed := timex.Since(start)

	ls := fmt.Sprintf(" - ip:%s - server:%s - req:%v - ret: %d ", peerInfo.Addr.String(), serverName, req, code)
	if err != nil {
		logx.WithContext(ctx).WithDuration(elapsed).Error("[rpc-client-log] fail" + ls+", - error: "+err.Error())
	} else {
		logx.WithContext(ctx).WithDuration(elapsed).Info("[rpc-client-log] ok" + ls)
	}

	if elapsed > slowThreshold {
		logx.WithContext(ctx).WithDuration(elapsed).Slow("[rpc-client-log] ok - slowcall " + ls)
	}

	return err
}
