package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mojtabaRKS/link_shortener/internal/bootstrap"
	"github.com/mojtabaRKS/link_shortener/internal/rest/middlewares"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Server *bootstrap.Server
	Router *mux.Router
}

func (app *App) Initilize() {
	// create http request handler
	app.Router = mux.NewRouter()
	// load routes
	app.Router.Use(middlewares.Logging)
	app.InitilizeRoutes()
}

func (app *App) RunServer() {

	portNo := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	log.Println(fmt.Sprintf("Starting HTTP server on port %s...", portNo))

	log.Fatal(http.ListenAndServe(portNo, app.Router))

}
