package builder

import (
	"github.com/mfaizudd/nodebat-go/validation"
	"github.com/mfaizudd/nodebat-go/validators"
)

type ValidationBuilder struct {
	validation *validation.Validation
	field      string
}

func New(v *validation.Validation, field string) *ValidationBuilder {
	return &ValidationBuilder{v, field}
}

func (v *ValidationBuilder) Required(value interface{}) *ValidationBuilder {
    v.validation.Add(v.field, validators.Required(value))
    return v
}

func (v *ValidationBuilder) Min(value int, min int) *ValidationBuilder {
    v.validation.Add(v.field, validators.Min(value, min))
    return v
}

func (v *ValidationBuilder) Max(value int, max int) *ValidationBuilder {
    v.validation.Add(v.field, validators.Max(value, max))
    return v
}

func (v *ValidationBuilder) Range(value int, min int, max int) *ValidationBuilder {
    v.validation.Add(v.field, validators.Range(value, min, max))
    return v
}

func (v *ValidationBuilder) MinLength(value string, length int) *ValidationBuilder {
    v.validation.Add(v.field, validators.MinLength(value, length))
    return v
}

func (v *ValidationBuilder) MaxLength(value string, length int) *ValidationBuilder {
    v.validation.Add(v.field, validators.MaxLength(value, length))
    return v
}

func (v *ValidationBuilder) Length(value string, min int, max int) *ValidationBuilder {
    v.validation.Add(v.field, validators.Length(value, min, max))
    return v
}

func (v *ValidationBuilder) IsEmail(value string) *ValidationBuilder {
    v.validation.Add(v.field, validators.IsEmail(value))
    return v
}

