package validation

import (
	"time"
)

func validateType[T any](value interface{}, builder *Builder) (T, bool) {
    v, ok := value.(T)
    if !ok {
        builder.add("Invalid type", "invalid_type")
    }
    return v, ok
}

type Builder struct {
	validation *Validation
	field      string
	value      interface{}
}

func NewBuilder(v *Validation, field string, value interface{}) *Builder {
	return &Builder{v, field, value}
}

func (v *Builder) add(message, tag string) {
    err := NewFieldError(v.field, message, tag, v.value)
    v.validation.AddError(v.field, err)
}

func (v *Builder) Required() *Builder {
	v.validation.Add(v.field, Required(v.value))
	return v
}

func (v *Builder) Min(min int) *Builder {
    value, ok := validateType[int](v.value, v)
    if ok {
        v.validation.Add(v.field, Min(value, min))
    }
    return v
}

func (v *Builder) Max(max int) *Builder {
    value, ok := validateType[int](v.value, v)
    if ok {
        v.validation.Add(v.field, Max(value, max))
    }
    return v
}

func (v *Builder) Range(min, max int) *Builder {
    value, ok := validateType[int](v.value, v)
    if ok {
        v.validation.Add(v.field, Range(value, min, max))
    }
    return v
}

func (v *Builder) MinLength(min int) *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, MinLength(value, min))
    }
    return v
}

func (v *Builder) MaxLength(max int) *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, MaxLength(value, max))
    }
    return v
}

func (v *Builder) Length(min, max int) *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, Length(value, min, max))
    }
    return v
}

func (v *Builder) OneOf(values ...string) *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, OneOf(value, values...))
    }
    return v
}

func (v *Builder) IsEmail() *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, IsEmail(value))
    }
    return v
}

func (v *Builder) IsAlphanumeric() *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, IsAlphanumeric(value))
    }
    return v
}

func (v *Builder) IsISO8601() *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, IsISO8601(value))
    }
    return v
}

func (v *Builder) IsISO8601Date() *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, IsISO8601Date(value))
    }
    return v
}

func (v *Builder) IsPhone() *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, IsPhone(value))
    }
    return v
}

func (v *Builder) IsUUID() *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, IsUUID(value))
    }
    return v
}

func (v *Builder) IsOnlyDigits() *Builder {
    value, ok := validateType[string](v.value, v)
    if ok {
        v.validation.Add(v.field, IsOnlyDigits(value))
    }
    return v
}

func (v *Builder) MinDate(minDate time.Time) *Builder {
    value, ok := validateType[time.Time](v.value, v)
    if ok {
        v.validation.Add(v.field, MinDate(value, minDate))
    }
    return v
}

func (v *Builder) MaxDate(maxDate time.Time) *Builder {
    value, ok := validateType[time.Time](v.value, v)
    if ok {
        v.validation.Add(v.field, MaxDate(value, maxDate))
    }
    return v
}

func (v *Builder) BetweenDate(minDate, maxDate time.Time) *Builder {
    value, ok := validateType[time.Time](v.value, v)
    if ok {
        v.validation.Add(v.field, BetweenDate(value, minDate, maxDate))
    }
    return v
}

func (v *Builder) Custom(validator Validator) *Builder {
    v.validation.Add(v.field, validator)
    return v
}
