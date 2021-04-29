package main

import (
	"flag"
	"fmt"
	"go-zero-study/app/service/comments/config"
	"go-zero-study/app/service/comments/internal/server/rpc"
	"go-zero-study/app/service/comments/internal/service"
	"go-zero-study/core/conf"
	"go-zero-study/zrpc"

	"go-zero-study/app/service/comments/api"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "app/service/comments/cmd/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := service.NewServiceContext(c)
	srv := rpc.NewCommentsRPCServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		api.RegisterCommentsRPCServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
