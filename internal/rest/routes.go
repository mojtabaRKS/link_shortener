package rest

import (
	"net/http"

	auth_v1 "github.com/mojtabaRKS/link_shortener/internal/rest/controllers/v1/auth"
)

func (app *App) InitilizeRoutes() {

	s := app.Router.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/register", auth_v1.Register).Methods(http.MethodPost)
}
