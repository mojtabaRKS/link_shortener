package auth

import (
	"fmt"
	"net/http"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,unique"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
	UserName string `json:"username" validate:"required,unique"`
}


func Register(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)
	// errs, request := validation.Validate(c, &RegisterRequest{})

	// if len(errs) != 0 {
	// 	util.RespondValidationError(c, errs)
	// 	return
	// }

	
}
