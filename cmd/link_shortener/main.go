package main

import (
	"github.com/mojtabaRKS/link_shortener/internal/rest"
	"github.com/mojtabaRKS/link_shortener/internal/bootstrap"
)

func main() {
	server := &bootstrap.Server{}

	// load REST server
	rest_app := &rest.App{Server: server}
	rest_app.Initilize()
	rest_app.RunServer()

	// load GRPC server
}
