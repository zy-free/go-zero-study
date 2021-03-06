package internal

import (
	"google.golang.org/grpc/keepalive"
	"net"
	"time"

	"go-zero-study/core/proc"
	"go-zero-study/core/stat"
	"go-zero-study/zrpc/internal/serverinterceptors"
	"google.golang.org/grpc"
)

type (
	ServerOption func(options *rpcServerOptions)

	rpcServerOptions struct {
		metrics *stat.Metrics
	}

	rpcServer struct {
		name string
		*baseRpcServer
	}
)

func init() {
	InitLogger()
}

func NewRpcServer(address string, opts ...ServerOption) Server {
	var options rpcServerOptions
	for _, opt := range opts {
		opt(&options)
	}
	if options.metrics == nil {
		options.metrics = stat.NewMetrics(address)
	}

	return &rpcServer{
		baseRpcServer: newBaseRpcServer(address, options.metrics),
	}
}

func (s *rpcServer) SetName(name string) {
	s.name = name
	s.baseRpcServer.SetName(name)
}

func (s *rpcServer) Start(register RegisterFn) error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		serverinterceptors.UnaryTracingInterceptor(s.name),
		serverinterceptors.UnaryRecoverInterceptor(),
		serverinterceptors.UnaryLogInterceptor(s.metrics),
		serverinterceptors.UnaryPrometheusInterceptor(),
		serverinterceptors.UnaryMetaDataInterceptor(),
	}
	unaryInterceptors = append(unaryInterceptors, s.unaryInterceptors...)
	streamInterceptors := []grpc.StreamServerInterceptor{
		//serverinterceptors.StreamCrashInterceptor,
	}
	streamInterceptors = append(streamInterceptors, s.streamInterceptors...)
	keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:                  time.Second*60,
		Timeout:               time.Second*10,
	})
	options := append(s.options,keepParam, WithUnaryServerInterceptors(unaryInterceptors...),
		WithStreamServerInterceptors(streamInterceptors...))
	server := grpc.NewServer(options...)
	register(server)
	// we need to make sure all others are wrapped up
	// so we do graceful stop at shutdown phase instead of wrap up phase
	shutdownCalled := proc.AddShutdownListener(func() {
		server.GracefulStop()
	})
	err = server.Serve(lis)
	shutdownCalled()

	return err
}

func WithMetrics(metrics *stat.Metrics) ServerOption {
	return func(options *rpcServerOptions) {
		options.metrics = metrics
	}
}
