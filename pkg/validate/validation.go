package validate

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/italorfeitosa/go-shorten/pkg/errs"
)

var validate = validator.New()

func Struct(v any) error {
	code := fmt.Sprintf("Invalid%s", getTypeName(v))

	validationError := errs.New(errs.Code(code), errs.Validation)

	err := validate.Struct(v)
	if err == nil {
		return nil
	}

	validatorValidationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		validationError.Reason = err.Error()
		return validationError
	}

	for i, err := range validatorValidationErrors {
		if i == 0 {
			validationError.Reason = err.Error()
			continue
		}

		validationError.Reason = fmt.Sprintf("%s\n%s", validationError.Reason, err.Error())
	}

	return validationError
}

func getTypeName(v any) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}
