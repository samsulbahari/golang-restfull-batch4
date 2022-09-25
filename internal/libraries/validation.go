package libraries

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validation(err error) []string {
	errorMessages := []string{}
	for _, e := range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("error on field %s , condition %s %s", e.Field(), e.Param(), e.Tag())
		errorMessages = append(errorMessages, errorMessage)
	}
	return errorMessages
}
