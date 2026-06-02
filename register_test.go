// register_test.go
package v10_extra_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	v10_extra "github.com/knightshrestha/v10_extra"
)

// mock ogen types
type OptString struct {
	Value string
	Set   bool
}

type OptInt64 struct {
	Value int64
	Set   bool
}

type registerStruct struct {
	Title string    `validate:"notblank,min=3"`
	Name  OptString `validate:"omitempty,min=3"`
	ID    OptInt64  `validate:"omitempty,min=1"`
}

func newValidator() *validator.Validate {
	v := validator.New()
	v10_extra.Register(v, OptString{}, OptInt64{})
	return v
}

func TestRegister_AllValid(t *testing.T) {
	v := newValidator()
	err := v.Struct(registerStruct{
		Title: "hello",
		Name:  OptString{Set: true, Value: "world"},
		ID:    OptInt64{Set: true, Value: 1},
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestRegister_OptFields_NotSet_Passes(t *testing.T) {
	v := newValidator()
	err := v.Struct(registerStruct{
		Title: "hello",
		Name:  OptString{Set: false},
		ID:    OptInt64{Set: false},
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestRegister_TitleBlank_Fails(t *testing.T) {
	v := newValidator()
	err := v.Struct(registerStruct{
		Title: "   ",
		Name:  OptString{Set: false},
		ID:    OptInt64{Set: false},
	})
	if err == nil {
		t.Error("expected error for blank title, got nil")
	}
}

func TestRegister_OptString_Set_Empty_Fails(t *testing.T) {
	v := newValidator()
	err := v.Struct(registerStruct{
		Title: "hello",
		Name:  OptString{Set: true, Value: ""},
		ID:    OptInt64{Set: false},
	})
	if err == nil {
		t.Error("expected error for empty OptString when set, got nil")
	}
}

func TestRegister_OptInt64_Set_Zero_Fails(t *testing.T) {
	v := newValidator()
	err := v.Struct(registerStruct{
		Title: "hello",
		Name:  OptString{Set: false},
		ID:    OptInt64{Set: true, Value: 0},
	})
	if err == nil {
		t.Error("expected error for zero OptInt64 when set, got nil")
	}
}