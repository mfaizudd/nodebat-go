package validators

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/mfaizudd/nodebat-go/validation"
)

// Required checks if the data is nil or empty string
func Required(data interface{}) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		msg := fmt.Sprintf("%s is required", field)
		if data == nil {
			return validation.NewFieldError(field, msg, "required", nil), false
		}
		if str, ok := data.(string); ok && str == "" {
			return validation.NewFieldError(field, msg, "required", str), false
		}
		return nil, true
	}
}

// Min checks if the data is at least min
//
// Has one parameter: min (int)
func Min(data int, min int) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		msg := fmt.Sprintf("%s must be at least %d", field, min)
		if data < min {
			return validation.NewFieldError(field, msg, "min", data), false
		}
		return nil, true
	}
}

// Max checks if the data is at most max
//
// Has one parameter: max (int)
func Max(data int, max int) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		msg := fmt.Sprintf("%s must be at most %d", field, max)
		if data > max {
			return validation.NewFieldError(field, msg, "max", data), false
		}
		return nil, true
	}
}

// Range checks if the data is between min and max
//
// Has two parameters: min (int) and max (int)
func Range(data int, min int, max int) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		msg := fmt.Sprintf("%s must be between %d and %d", field, min, max)
		if data < min || data > max {
			return validation.NewFieldError(field, msg, "range", data), false
		}
		return nil, true
	}
}

// MinLength checks if the data is at least min characters long
//
// Has one parameter: min (int)
func MinLength(data string, min int) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		msg := fmt.Sprintf("%s must be at least %d characters long", field, min)
		if len(data) < min {
			return validation.NewFieldError(field, msg, "min_length", data), false
		}
		return nil, true
	}
}

// MaxLength checks if the data is at most max characters long
//
// Has one parameter: max (int)
func MaxLength(data string, max int) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		msg := fmt.Sprintf("%s must be at most %d characters long", field, max)
		if len(data) > max {
			return validation.NewFieldError(field, msg, "max_length", data), false
		}
		return nil, true
	}
}

// OneOf checks if the data is in the collection
//
// Has one parameter named "collection" which is a slice of strings
func OneOf(item string, collection ...string) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		for _, it := range collection {
			if item == it {
				return nil, true
			}
		}
		msg := fmt.Sprintf("%s is not in the collection", field)
		err = validation.NewFieldError(field, msg, "one_of", item)
		err.SetParam("collection", collection)
		return err, false
	}
}

// IsEmail checks if the data is a valid email address
func IsEmail(email string) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		_, emailErr := mail.ParseAddress(email)
		if emailErr != nil {
			msg := fmt.Sprintf("%s is not a valid email address", field)
			return validation.NewFieldError(field, msg, "is_email", email), false
		}
		return nil, true
	}
}

// IsAlphanumeric checks if the data is alphanumeric excluding space
func IsAlphanumeric(value string) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		if !regexp.MustCompile(`^([a-zA-Z0-9])+$`).Match([]byte(value)) {
			msg := fmt.Sprintf("%s must be alphanumeric", field)
			return validation.NewFieldError(field, msg, "is_alphanumeric", value), false
		}
		return nil, true
	}
}
