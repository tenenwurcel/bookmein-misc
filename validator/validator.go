package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validator() *validator.Validate {
	return validate
}

func ValidateIntId(id int) error {
	return validate.Var(id, "gt=0")
}

func ValidateUuid(id string) error {
	return validate.Var(id, "uuid")
}
