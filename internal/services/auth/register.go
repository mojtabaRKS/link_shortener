package auth

import (
	"github.com/mojtabaRKS/link_shortener/internal/entities"
	"github.com/mojtabaRKS/link_shortener/pkg/db"
)

func Register(name, username, email, password string) (string, error) {

	user := entities.User{
		Name:     name,
		Email:    email,
		Username: username,
	}

	db.GetInstance().Create(&user)

	return "user.Email", nil
}

func encryptPassword(pswd string) (string, error) {
	return "encrypted password",nil
}
