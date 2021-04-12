package main

import (
	"flag"
	"fmt"

	"go-zero-study/app/api/member/config"
	"go-zero-study/app/api/member/internal/server/http"
	"go-zero-study/app/api/member/internal/service"

	"go-zero-study/core/conf"
	"go-zero-study/rest"
)

var configFile = flag.String("f", "./app/api/member/cmd/member-api.yaml", "the config file")

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
