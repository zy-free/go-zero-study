package serverinterceptors

import (
	"context"
	"fmt"
	"go-zero-study/core/ecode"
	"go-zero-study/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"runtime"
)


func UnaryRecoverInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if rerr := recover(); rerr != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				rs := runtime.Stack(buf, false)
				if rs > size {
					rs = size
				}
				buf = buf[:rs]
				pl := fmt.Sprintf("grpc server panic: %v\n%v\n%s\n", req, rerr, buf)
				fmt.Fprint(os.Stderr, pl)
				logx.WithContext(ctx).Error(pl)
				err = status.Errorf(codes.Unknown, ecode.ServerErr.Error())
			}
		}()
		resp, err = handler(ctx, req)
		return
	}
}
