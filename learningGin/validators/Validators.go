package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "cool")
}
