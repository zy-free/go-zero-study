package serverinterceptors

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-study/core/ecode"
	"time"

	"go-zero-study/core/logx"
	"go-zero-study/core/stat"
	"go-zero-study/core/timex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const serverSlowThreshold = time.Millisecond * 500

func UnaryLogInterceptor(metrics *stat.Metrics) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		startTime := timex.Now()
		var remoteIP string
		if peerInfo, ok := peer.FromContext(ctx); ok {
			remoteIP = peerInfo.Addr.String()
		}
		var quota float64
		if deadline, ok := ctx.Deadline(); ok {
			quota = time.Until(deadline).Seconds()
		}

		// call server handler
		resp, err = handler(ctx, req)
		// after server response
		code := ecode.Cause(err).Code()
		elapsed := timex.Since(startTime)
		content, _ := json.Marshal(req)

		ls := fmt.Sprintf(" - ip:%s - path:%s - req:%s - ret: %d - timeout_quota: %v", remoteIP, info.FullMethod, string(content), code, quota)
		if err != nil {
			logx.WithContext(ctx).WithDuration(elapsed).Error("[grpc-access-log] fail" + ls + ", - error: " + err.Error())
		} else {
			logx.WithContext(ctx).WithDuration(elapsed).Info("[grpc-access-log] ok" + ls)
		}

		if elapsed > serverSlowThreshold {
			logx.WithContext(ctx).WithDuration(elapsed).Slow("[grpc-access-log] ok - slowcall " + ls)
		}

		return
	}
}
