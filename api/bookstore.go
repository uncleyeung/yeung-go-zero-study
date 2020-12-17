package main

import (
	"flag"
	"fmt"

	"github.com/uncleyeung/yeung-go-zero-study/api/internal/config"
	"github.com/uncleyeung/yeung-go-zero-study/api/internal/handler"
	"github.com/uncleyeung/yeung-go-zero-study/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "api/etc/yeung-go-zero-study-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
