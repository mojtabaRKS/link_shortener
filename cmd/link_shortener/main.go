package main

import (
	"github.com/mojtabaRKS/link_shortener/internal/webserver"
)

func main() {
	app := webserver.App{}

	app.Initilize()
	app.RunServer()
}
