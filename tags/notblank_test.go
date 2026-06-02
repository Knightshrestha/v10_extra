// tags/notblank_test.go
package tags_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/knightshrestha/v10_extra/tags"
)

type notBlankStruct struct {
	Title string `validate:"notblank"`
}

func newValidator() *validator.Validate {
	v := validator.New()
	tags.RegisterNotBlank(v)
	return v
}

func TestNotBlank_Valid(t *testing.T) {
	v := newValidator()
	err := v.Struct(notBlankStruct{Title: "hello"})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestNotBlank_Empty(t *testing.T) {
	v := newValidator()
	err := v.Struct(notBlankStruct{Title: ""})
	if err == nil {
		t.Error("expected error for empty string, got nil")
	}
}

func TestNotBlank_WhitespaceOnly(t *testing.T) {
	v := newValidator()
	err := v.Struct(notBlankStruct{Title: "   "})
	if err == nil {
		t.Error("expected error for whitespace-only string, got nil")
	}
}

func TestNotBlank_WhitespaceAround(t *testing.T) {
	v := newValidator()
	err := v.Struct(notBlankStruct{Title: "  hello  "})
	if err != nil {
		t.Errorf("expected nil for string with surrounding whitespace, got %v", err)
	}
}