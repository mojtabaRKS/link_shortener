package rest

import (
	"github.com/mojtabaRKS/link_shortener/internal/entities"
)

func (app *App) Migrate() {
	app.DB.AutoMigrate(
		&entities.User{},
	)
}
