// tags/notblank.go
package tags

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// notblank validates that a string is not empty or whitespace-only.
func notblank(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return strings.TrimSpace(field.String()) != ""
	default:
		return true
	}
}

// RegisterNotBlank registers the "notblank" tag with the validator.
func RegisterNotBlank(v *validator.Validate) {
	v.RegisterValidation("notblank", notblank)
}