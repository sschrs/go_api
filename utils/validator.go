package utils

import (
	"data_app/api"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(s any) (bool, api.ErrorValidation) {

	var errorValidation api.ErrorValidation

	if err := validate.Struct(s); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			for _, err := range err.(validator.ValidationErrors) {
				errorValidation.Field = err.Field()
				errorValidation.Message = err.Error()
				errorValidation.Tag = err.Tag()
			}
		}
		return false, errorValidation
	}
	return true, errorValidation
}
