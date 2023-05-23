package validation

import "time"

type Builder struct {
	validation *Validation
	field      string
	value      interface{}
}

// NewBuilder creates a new Builder
func NewBuilder(v *Validation, field string, value interface{}) *Builder {
	return &Builder{v, field, value}
}

// add adds an error to the validation object
// Required checks if the data is nil or empty string
func (v *Builder) Required() *Builder {
	if v.hasError() {
		return v
	}
	v.validation.Add(v.field, Required(v.value))
	return v
}

// Min checks if the data is at least min
//
// Has one parameter: min (int64)
func (v *Builder) Min(min int64) *Builder {
	return v.MinInt(min)
}

// Max checks if the data is at most max
//
// Has one parameter: max (int64)
func (v *Builder) Max(max int64) *Builder {
	return v.MaxInt(max)
}

// Range checks if the data is between min and max
//
// Has two parameters: min and max (int64)
func (v *Builder) Range(min, max int64) *Builder {
	return v.RangeInt(min, max)
}

// MinInt checks if the data is at least min
//
// Has one parameter: min (int64)
func (v *Builder) MinInt(min int64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getInt()
	if ok {
		v.validation.Add(v.field, Min(value, min))
	}
	return v
}

// MaxInt checks if the data is at most max
//
// Has one parameter: max (int64)
func (v *Builder) MaxInt(max int64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getInt()
	if ok {
		v.validation.Add(v.field, Max(value, max))
	}
	return v
}

// RangeInt checks if the data is between min and max
//
// Has two parameters: min and max (int64)
func (v *Builder) RangeInt(min, max int64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getInt()
	if ok {
		v.validation.Add(v.field, Range(value, min, max))
	}
	return v
}

// MinUint checks if the data is at least min
//
// Has one parameter: min (uint64)
func (v *Builder) MinUint(min uint64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getUint()
	if ok {
		v.validation.Add(v.field, Min(value, min))
	}
	return v
}

// MaxUint checks if the data is at most max
//
// Has one parameter: max (uint64)
func (v *Builder) MaxUint(max uint64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getUint()
	if ok {
		v.validation.Add(v.field, Max(value, max))
	}
	return v
}

// RangeUint checks if the data is between min and max
//
// Has two parameters: min and max (uint64)
func (v *Builder) RangeUint(min, max uint64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getUint()
	if ok {
		v.validation.Add(v.field, Range(value, min, max))
	}
	return v
}

// MinFloat checks if the data is at least min
//
// Has one parameter: min (float64)
func (v *Builder) MinFloat(min float64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getFloat()
	if ok {
		v.validation.Add(v.field, Min(value, min))
	}
	return v
}

// MaxFloat checks if the data is at most max
//
// Has one parameter: max (float64)
func (v *Builder) MaxFloat(max float64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getFloat()
	if ok {
		v.validation.Add(v.field, Max(value, max))
	}
	return v
}

// RangeFloat checks if the data is between min and max
//
// Has two parameters: min and max (float64)
func (v *Builder) RangeFloat(min, max float64) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getFloat()
	if ok {
		v.validation.Add(v.field, Range(value, min, max))
	}
	return v
}

// MinLength checks if the data is at least min characters long
//
// Has one parameter: min (int)
func (v *Builder) MinLength(min int) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, MinLength(value, min))
	}
	return v
}

// MaxLength checks if the data is at most max characters long
//
// Has one parameter: max (int)
func (v *Builder) MaxLength(max int) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, MaxLength(value, max))
	}
	return v
}

// Length checks if the data is between min and max characters long
//
// Has one parameter: length (int)
func (v *Builder) Length(min, max int) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, Length(value, min, max))
	}
	return v
}

// OneOf checks if the data is in the collection
//
// Has one parameter named "collection" which is a slice of strings
func (v *Builder) OneOf(values ...string) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, OneOf(value, values...))
	}
	return v
}

// IsEmail checks if the data is a valid email address
func (v *Builder) IsEmail() *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, IsEmail(value))
	}
	return v
}

// IsAlphanumeric checks if the data is alphanumeric excluding space
func (v *Builder) IsAlphanumeric() *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, IsAlphanumeric(value))
	}
	return v
}

// IsISO8601 checks if the data is a valid ISO8601 date
func (v *Builder) IsISO8601() *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, IsISO8601(value))
	}
	return v
}

// IsISO8601Date checks if the data is a valid ISO8601 date
func (v *Builder) IsISO8601Date() *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, IsISO8601Date(value))
	}
	return v
}

// IsPhone checks if the data is a valid phone number
func (v *Builder) IsPhone() *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, IsPhone(value))
	}
	return v
}

// IsUUID checks if the data is a valid UUID
func (v *Builder) IsUUID() *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, IsUUID(value))
	}
	return v
}

// IsOnlyDigits checks if the data contains only digits
func (v *Builder) IsOnlyDigits() *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getString()
	if ok {
		v.validation.Add(v.field, IsOnlyDigits(value))
	}
	return v
}

// MinDate checks if the date is after the given date
//
// Has one parameter: minDate (time.Time)
func (v *Builder) MinDate(minDate time.Time) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getTime()
	if ok {
		v.validation.Add(v.field, MinDate(value, minDate))
	}
	return v
}

// MaxDate checks if the date is before the given date
//
// Has one parameter: maxDate (time.Time)
func (v *Builder) MaxDate(maxDate time.Time) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getTime()
	if ok {
		v.validation.Add(v.field, MaxDate(value, maxDate))
	}
	return v
}

// BetweenDate checks if the date is between the given dates
//
// Has two parameters: minDate (time.Time), maxDate (time.Time)
func (v *Builder) BetweenDate(minDate, maxDate time.Time) *Builder {
	if v.hasError() {
		return v
	}
	value, ok := v.getTime()
	if ok {
		v.validation.Add(v.field, BetweenDate(value, minDate, maxDate))
	}
	return v
}

// Custom adds a custom validator to the validation
func (v *Builder) Custom(validator Validator) *Builder {
	if v.hasError() {
		return v
	}
	v.validation.Add(v.field, validator)
	return v
}

// MinCount checks if the slice/array/map has a minimum number of elements
//
// Has one parameter: min (int)
func (v *Builder) MinCount(min int) *Builder {
	if v.hasError() {
		return v
	}
	v.validation.Add(v.field, MinCount(v.value, min))
	return v
}

// MaxCount checks if the slice/array/map has a maximum number of elements
//
// Has one parameter: max (int)
func (v *Builder) MaxCount(max int) *Builder {
	if v.hasError() {
		return v
	}
	v.validation.Add(v.field, MaxCount(v.value, max))
	return v
}

// Numeric checks if the value is a number
func (v *Builder) Numeric() *Builder {
	if v.hasError() {
		return v
	}
	v.validation.Add(v.field, Numeric(v.value))
	return v
}
