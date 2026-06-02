// register.go
package v10_extra

import (
	"github.com/go-playground/validator/v10"
	"github.com/knightshrestha/v10_extra/ogenval"
	"github.com/knightshrestha/v10_extra/tags"
)

// Register registers all custom type funcs and validation tags.
// Pass ogen Opt* types you want supported, e.g.:
//
//	v10_extra.Register(validate, api.OptString{}, api.OptInt64{})
func Register(v *validator.Validate, ogenTypes ...interface{}) {
	tags.RegisterNotBlank(v)

	if len(ogenTypes) > 0 {
		ogenval.RegisterOgenTypes(v, ogenTypes...)
	}
}