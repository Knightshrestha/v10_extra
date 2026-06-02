# v10_extra

Extensions for [go-playground/validator/v10](https://github.com/go-playground/validator) — adds support for [ogen](https://github.com/ogen-go/ogen) generated `Opt*` types and extra validation tags.

## Installation

```bash
go get github.com/knightshrestha/v10_extra
```

## Usage

```go
import (
    "github.com/go-playground/validator/v10"
    v10_extra "github.com/knightshrestha/v10_extra"
)

validate := validator.New()

// pass any ogen Opt* types your project uses
v10_extra.Register(validate, api.OptString{}, api.OptBool{}, api.OptInt{}, api.OptInt64{})
```

## Features

### Ogen `Opt*` type support

Ogen generates optional fields as structs like:

```go
type OptString struct {
    Value string
    Set   bool
}
```

By default, `go-playground/validator` cannot handle these — standard tags like `min`, `max`, `email` won't work correctly.

After calling `v10_extra.Register`, ogen `Opt*` types behave like their native counterparts:

```go
// plain string
Username string `validate:"required,min=3,max=20"`

// ogen OptString — behaves identically, skips validation if not set
Username OptString `validate:"omitempty,min=3,max=20"`
```

To avoid writing validate tags by hand, use ogen's `x-oapi-codegen-extra-tags` extension in your OpenAPI spec:

```yaml
components:
  schemas:
    UpdateUserRequest:
      type: object
      properties:
        username:
          type: string
          x-oapi-codegen-extra-tags:
            validate: "omitempty,min=3,max=20"
        age:
          type: integer
          format: int64
          x-oapi-codegen-extra-tags:
            validate: "omitempty,min=1"
```

Ogen will generate the struct tags automatically:

```go
type UpdateUserRequest struct {
    Username OptString `json:"username" validate:"omitempty,min=3,max=20"`
    Age      OptInt64  `json:"age"      validate:"omitempty,min=1"`
}
```

> **Note:** Always use `omitempty` with `Opt*` fields. The "skip if not set" behavior requires it.

### `notblank` tag

`go-playground/validator`'s `min=1` passes whitespace-only strings like `"   "`. The `notblank` tag fails if the string is empty or whitespace-only.

```go
Title string `validate:"required,notblank,min=3"`
```

| Value | Result |
|-------|--------|
| `"hello"` | ✓ pass |
| `""` | ✗ fail |
| `"   "` | ✗ fail |
| `"  hello  "` | ✓ pass |

> **Note:** `notblank` only checks — it does not trim. Sanitization should be done explicitly before validation: `title = strings.TrimSpace(title)`

## Supported `Opt*` types

Pass any ogen generated `Opt*` type to `Register`. Tested with:

- `OptString`
- `OptBool`
- `OptInt`
- `OptInt64`

Any other ogen `Opt*` type with `Value` and `Set` fields will also work.

## Credits

Library authored with major assistance from [Claude](https://claude.ai) (Anthropic).