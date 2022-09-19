package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mojtabaRKS/link_shortener/pkg/db"
	"github.com/mojtabaRKS/link_shortener/pkg/env"
	"github.com/mojtabaRKS/link_shortener/internal/rest/middlewares"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func init() {
	// first we load environment variables
	env.Load()
}

func (app *App) Initilize() {
	// create database instance
	connection := db.Connection{}
	app.DB = connection.Load().Conn
	// create database schema
	app.Migrate()
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
