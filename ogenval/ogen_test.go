// ogenval/ogen_test.go
package ogenval_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/knightshrestha/v10_extra/ogenval"
)

// local mocks
type OptString struct {
	Value string
	Set   bool
}

type OptBool struct {
	Value bool
	Set   bool
}

type OptInt struct {
	Value int
	Set   bool
}

type OptInt64 struct {
	Value int64
	Set   bool
}

func newValidator(types ...interface{}) *validator.Validate {
	v := validator.New()
	ogenval.RegisterOgenTypes(v, types...)
	return v
}

// --- OptString ---

type stringStruct struct {
	Name OptString `validate:"omitempty,min=3,max=10"`
}

func TestOptString_NotSet_Passes(t *testing.T) {
	v := newValidator(OptString{})
	err := v.Struct(stringStruct{Name: OptString{Set: false, Value: ""}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestOptString_Set_Valid(t *testing.T) {
	v := newValidator(OptString{})
	err := v.Struct(stringStruct{Name: OptString{Set: true, Value: "hello"}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestOptString_Set_TooShort(t *testing.T) {
	v := newValidator(OptString{})
	err := v.Struct(stringStruct{Name: OptString{Set: true, Value: "hi"}})
	if err == nil {
		t.Error("expected error for too short string, got nil")
	}
}

func TestOptString_Set_TooLong(t *testing.T) {
	v := newValidator(OptString{})
	err := v.Struct(stringStruct{Name: OptString{Set: true, Value: "hello_world_123"}})
	if err == nil {
		t.Error("expected error for too long string, got nil")
	}
}

func TestOptString_Set_Empty(t *testing.T) {
	v := newValidator(OptString{})
	err := v.Struct(stringStruct{Name: OptString{Set: true, Value: ""}})
	if err == nil {
		t.Error("expected error for empty string when set, got nil")
	}
}

// --- OptBool ---

type boolStruct struct {
	Active OptBool `validate:"omitempty"`
}

func TestOptBool_NotSet_Passes(t *testing.T) {
	v := newValidator(OptBool{})
	err := v.Struct(boolStruct{Active: OptBool{Set: false}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestOptBool_Set_Passes(t *testing.T) {
	v := newValidator(OptBool{})
	err := v.Struct(boolStruct{Active: OptBool{Set: true, Value: true}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

// --- OptInt ---

type intStruct struct {
	Age OptInt `validate:"omitempty,min=1,max=120"`
}

func TestOptInt_NotSet_Passes(t *testing.T) {
	v := newValidator(OptInt{})
	err := v.Struct(intStruct{Age: OptInt{Set: false}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestOptInt_Set_Valid(t *testing.T) {
	v := newValidator(OptInt{})
	err := v.Struct(intStruct{Age: OptInt{Set: true, Value: 25}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestOptInt_Set_OutOfRange(t *testing.T) {
	v := newValidator(OptInt{})
	err := v.Struct(intStruct{Age: OptInt{Set: true, Value: 200}})
	if err == nil {
		t.Error("expected error for out of range int, got nil")
	}
}

// --- OptInt64 ---

type int64Struct struct {
	ID OptInt64 `validate:"omitempty,min=1"`
}

func TestOptInt64_NotSet_Passes(t *testing.T) {
	v := newValidator(OptInt64{})
	err := v.Struct(int64Struct{ID: OptInt64{Set: false}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestOptInt64_Set_Valid(t *testing.T) {
	v := newValidator(OptInt64{})
	err := v.Struct(int64Struct{ID: OptInt64{Set: true, Value: 42}})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestOptInt64_Set_Invalid(t *testing.T) {
	v := newValidator(OptInt64{})
	err := v.Struct(int64Struct{ID: OptInt64{Set: true, Value: 0}})
	if err == nil {
		t.Error("expected error for zero value, got nil")
	}
}