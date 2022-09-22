package validation

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var errorBag = make(map[string]string)

func Validate(request *http.Request, s interface{}) (map[string]string, interface{}) {

	validate = validator.New()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("could not read body: %s\n", err)
	}

	_ = json.Unmarshal(body, &s)

	err = validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorBag[err.Field()] = Map(strings.ToLower(err.Field()), err.Tag())
		}
	}

	return errorBag, s
}
