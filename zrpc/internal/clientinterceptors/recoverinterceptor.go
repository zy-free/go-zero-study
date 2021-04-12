package clientinterceptors

import (
	"context"
	"fmt"
	"go-zero-study/core/ecode"
	"go-zero-study/core/logx"
	"google.golang.org/grpc"
	"os"
	"runtime"
)


func RecoveryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			rs := runtime.Stack(buf, false)
			if rs > size {
				rs = size
			}
			buf = buf[:rs]
			pl := fmt.Sprintf("grpc client panic: %v\n%v\n%v\n%s\n", req, reply, rerr, buf)
			fmt.Fprintf(os.Stderr, pl)
			logx.WithContext(ctx).Error(pl)
			err = ecode.ServerErr
		}
	}()
	err = invoker(ctx, method, req, reply, cc, opts...)
	return

}
