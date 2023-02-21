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

// MinLength checks if the data is at least min characters long
func MinLength(data string, min int) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		msg := fmt.Sprintf("%s must be at least %d characters long", field, min)
		if len(data) < min {
			return validation.NewFieldError(field, msg, "min_length", data), false
		}
		return nil, true
	}
}

// OneOf checks if the data is in the collection
func OneOf(item string, collection ...string) func(field string) (err *validation.FieldError, ok bool) {
	return func(field string) (err *validation.FieldError, ok bool) {
		for _, it := range collection {
			if item == it {
				return nil, true
			}
		}
		msg := fmt.Sprintf("%s is not in the collection", field)
		return validation.NewFieldError(field, msg, "one_of", item), false
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
