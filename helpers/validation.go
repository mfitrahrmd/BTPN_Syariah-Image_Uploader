package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// CustomValidationErrorMessage customize error binding message
func CustomValidationErrorMessage(vErr validator.ValidationErrors) map[string]any {
	errs := make(map[string]any)

	for _, ve := range vErr {
		switch ve.Tag() {
		case "required":
			errs[ve.Field()] = fmt.Sprintf("%s is required", ve.Field())
		case "email":
			errs[ve.Field()] = fmt.Sprintf("%s must be valid email", ve.Field())
		case "min":
			errs[ve.Field()] = fmt.Sprintf("%s must be longer than or equal %s", ve.Field(), ve.Param())
		case "max":
			errs[ve.Field()] = fmt.Sprintf("%s must be less than or equal %s", ve.Field(), ve.Param())
		case "url":
			errs[ve.Field()] = fmt.Sprintf("%s must be valid url", ve.Field())
		}
	}

	return errs
}
