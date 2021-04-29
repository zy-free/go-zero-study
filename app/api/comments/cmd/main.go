package main

import (
	"flag"
	"fmt"
	"go-zero-study/app/api/comments/internal/server/http"
	"go-zero-study/app/api/comments/internal/service"
	"go-zero-study/core/conf"
	"go-zero-study/rest"

	"go-zero-study/app/api/comments/config"
)

var configFile = flag.String("f", "app/api/comments/cmd/comment-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := service.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	http.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
