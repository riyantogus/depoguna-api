package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)

type ValidationUtil interface {
	Validate(i interface{}) error
	ErrorMessage(err validator.FieldError) string
}

type validationUtil struct {
	validator *validator.Validate
}

func NewValidationUtil() ValidationUtil {
	return &validationUtil{
		validator: validator.New(),
	}
}

func (v *validationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func (v *validationUtil) ErrorMessage(fieldErr validator.FieldError) string {
	var msg string

	switch fieldErr.Tag() {
	case "required":
		msg = fmt.Sprintf("the %s field is required", fieldErr.Field())
	case "email":
		msg = fmt.Sprintf("the %s must be a valid email address", fieldErr.Field())
	case "min":
		msg = fmt.Sprintf("%s value must be greater than %s", fieldErr.Field(), fieldErr.Param())
	case "max":
		msg = fmt.Sprintf("%s value must be lower than %s", fieldErr.Field(), fieldErr.Param())
	default:
		msg = fmt.Sprintf("something is wrong with %s field", fieldErr.Field())
	}
	return msg
}
