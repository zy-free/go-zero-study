package internal

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go-zero-study/zrpc/internal/balancer/p2c"
	"go-zero-study/zrpc/internal/clientinterceptors"
	"go-zero-study/zrpc/internal/resolver"
	"google.golang.org/grpc"
)

const (
	dialTimeout = time.Second * 3
	separator   = '/'
)

func init() {
	resolver.RegisterResolver()
}

type (
	ClientOptions struct {
		Timeout     time.Duration
		DialOptions []grpc.DialOption
	}

	ClientOption func(options *ClientOptions)

	client struct {
		conn         *grpc.ClientConn
		interceptors []grpc.UnaryClientInterceptor
	}
)

func NewClient(target string, opts ...ClientOption) (*client, error) {
	var cli client
	opts = append(opts, WithDialOption(grpc.WithBalancerName(p2c.Name)))
	if err := cli.dial(target, opts...); err != nil {
		return nil, err
	}

	return &cli, nil
}

func (c *client) AddInterceptor(interceptor grpc.UnaryClientInterceptor) {
	c.interceptors = append(c.interceptors, interceptor)
}

func (c *client) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *client) buildDialOptions(opts ...ClientOption) []grpc.DialOption {
	var clientOptions ClientOptions
	for _, opt := range opts {
		opt(&clientOptions)
	}

	options := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		WithUnaryClientInterceptors(
			clientinterceptors.GrpcErrorInterceptor,
			clientinterceptors.MetaDataInterceptor,
			clientinterceptors.TracingInterceptor,
			clientinterceptors.LogInterceptor,
			clientinterceptors.RecoveryInterceptor,
			clientinterceptors.BreakerInterceptor,
			clientinterceptors.PrometheusInterceptor,
			clientinterceptors.TimeoutInterceptor(clientOptions.Timeout),
		),
	}
	for _, interceptor := range c.interceptors {
		options = append(options, WithUnaryClientInterceptors(interceptor))
	}

	return append(options, clientOptions.DialOptions...)
}

func (c *client) dial(server string, opts ...ClientOption) error {
	options := c.buildDialOptions(opts...)
	timeCtx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()
	conn, err := grpc.DialContext(timeCtx, server, options...)
	if err != nil {
		service := server
		if errors.Is(err, context.DeadlineExceeded) {
			pos := strings.LastIndexByte(server, separator)
			// len(server) - 1 is the index of last char
			if 0 < pos && pos < len(server)-1 {
				service = server[pos+1:]
			}
		}
		return fmt.Errorf("rpc dial: %s, error: %s, make sure rpc service %q is alread started",
			server, err.Error(), service)
	}

	c.conn = conn
	return nil
}

func WithDialOption(opt grpc.DialOption) ClientOption {
	return func(options *ClientOptions) {
		options.DialOptions = append(options.DialOptions, opt)
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(options *ClientOptions) {
		options.Timeout = timeout
	}
}
