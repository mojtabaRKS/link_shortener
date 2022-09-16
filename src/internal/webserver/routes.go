package webserver

import (
	"net/http"

	auth_v1 "github.com/mojtabaRKS/link_shortener/internal/webserver/controllers/v1/auth"
)

func (app *App) InitilizeRoutes() {
	r := app.Router.PathPrefix("/api").Subrouter() 
	
	r.PathPrefix("/v1").Subrouter()
	{
		r.HandleFunc("register", auth_v1.Register).Methods(http.MethodPost)
	}
}
