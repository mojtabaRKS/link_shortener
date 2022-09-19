package main

import (
	"github.com/mojtabaRKS/link_shortener/internal/rest"
)

func main() {
	app := rest.App{}

	app.Initilize()
	app.RunServer()
}
