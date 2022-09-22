package bootstrap

import (
	"github.com/mojtabaRKS/link_shortener/internal/entities"
	"github.com/mojtabaRKS/link_shortener/pkg/db"
	"gorm.io/gorm"

	"github.com/mojtabaRKS/link_shortener/pkg/env"
)

type Server struct {
	DB     *gorm.DB
}

func init() {
	// first we load environment variables
	env.Load()
}

func (s *Server) Boot() {
	s.DB = db.GetInstance()
	// create database schema
	s.Migrate()
}

func (s *Server) Migrate() {
	s.DB.AutoMigrate(
		&entities.User{},
	)
}
