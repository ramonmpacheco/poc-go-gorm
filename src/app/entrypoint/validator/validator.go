package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) ([]string, error) {
	structValidator := validator.New()
	err := structValidator.Struct(s)

	errs := make([]string, 0)

	if err == nil {
		return errs, nil
	}

	validationErrors := err.(validator.ValidationErrors)
	for _, v := range validationErrors {
		switch v.Tag() {
		case "required":
			errs = append(
				errs,
				fmt.Sprintf("%s is required", v.Field()))
		case "min":
			errs = append(
				errs,
				fmt.Sprintf("%s is required with min: %s", v.Field(), v.Param()))
		default:
			errs = append(
				errs,
				fmt.Sprintf("%s is invalid", v.Field()))
		}
	}
	return errs, errors.New("invalid data")
}
