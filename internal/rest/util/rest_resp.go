package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MainStructure struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func RespondError(w http.ResponseWriter, err error) {
	data := MainStructure{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		},
		Data: nil,
	}

	marshalled, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, string(marshalled))
}

func RespondValidationError(w http.ResponseWriter, errs map[string]string) {
	data := MainStructure{
		Meta: Meta{
			Code:    http.StatusUnprocessableEntity,
			Message: "validation error",
			Success: false,
		},
		Data: errs,
	}

	marshalled, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	fmt.Fprintln(w, string(marshalled))
}
