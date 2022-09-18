package entities

import (
	// "html"
	// "strings"
	// "github.com/mojtabaRKS/link_shortener/pkg/db"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}


// func (u *User) Save() (*User, error) {
// 	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)

// 	// if err != nil {
// 	// 	return u, err
// 	// }

// 	// u.Password = string(hashedPassword)
// 	// u.Username = html.EscapeString(strings.TrimSpace(u.Username))

// 	// res, err := db.
// }