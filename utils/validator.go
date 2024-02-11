package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) ([]string, error) {
	var errs []string
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return nil, err
		}

		for _, err := range err.(validator.ValidationErrors) {
			var element string
			element = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag())
			errs = append(errs, element)
		}
		return errs, nil
	}
	return nil, nil
}
