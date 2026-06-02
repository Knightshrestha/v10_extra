// ogenval/ogen.go
package ogenval

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// optValue extracts the underlying value from ogen Opt* types.
// Returns nil if Set is false, otherwise returns the Value field.
func optValue(field reflect.Value) interface{} {
	if field.Kind() != reflect.Struct {
		return nil
	}

	setField := field.FieldByName("Set")
	valueField := field.FieldByName("Value")

	if !setField.IsValid() || !valueField.IsValid() {
		return nil
	}

	if !setField.Bool() {
		return nil
	}

	// wrap in pointer so omitempty doesn't treat zero values as empty
	ptr := reflect.New(valueField.Type())
	ptr.Elem().Set(valueField)
	return ptr.Interface()
}

// RegisterOgenTypes registers RegisterCustomTypeFunc for common ogen Opt* types.
// Supports: OptString, OptBool, OptInt, OptInt64 and their OptNil* variants.
func RegisterOgenTypes(v *validator.Validate, types ...interface{}) {
	v.RegisterCustomTypeFunc(optValue, types...)
}
