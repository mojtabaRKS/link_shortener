package webserver

import (
	"fmt"
	"net/http"

	auth_v1 "github.com/mojtabaRKS/link_shortener/internal/webserver/controllers/v1/auth"
)

func (app *App) InitilizeRoutes() {
	s := app.Router.PathPrefix("/api").Subrouter() 
	
	s = s.PathPrefix("/v1").Subrouter()

	s.HandleFunc("register", auth_v1.Register).Methods(http.MethodPost)

	fmt.Print(app.Router)
}
