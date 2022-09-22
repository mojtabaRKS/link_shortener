package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mojtabaRKS/link_shortener/internal/rest/util"
	AuthService "github.com/mojtabaRKS/link_shortener/internal/services/auth"
	"github.com/mojtabaRKS/link_shortener/pkg/validation"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
	UserName string `json:"username" validate:"required"`
}

type TokenStruct struct{
	Token string `json:"token"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	req := RegisterRequest{}
	errs, _ := validation.Validate(r, &req)

	if len(errs) != 0 {
		util.RespondValidationError(w, errs)
		return
	}

	token, err := AuthService.Register(
		req.Name,
		req.UserName,
		req.Email,
		req.Password,
	)

	if err != nil {
		util.RespondError(w, err)
		return
	}

	res := util.MainStructure{
		Meta: util.Meta{
			Success: true,
			Code: http.StatusCreated,
			Message: "user registered successfully !",
		},
		Data: TokenStruct{Token: token},
	}

	marshalled, err := json.Marshal(res)

	if err != nil {
		util.RespondError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(marshalled))
	return
}
