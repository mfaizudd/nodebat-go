package validators

import (
	"fmt"
	"net/mail"
	"regexp"
)

// Required checks if the data is nil or empty string
func Required(data interface{}) func(string) (string, bool) {
	return func(s string) (string, bool) {
		msg := fmt.Sprintf("%s is required", s)
		if data == nil {
			return msg, false
		}
		if str, ok := data.(string); ok && str == "" {
			return msg, false
		}
		return "", true
	}
}

// MinLength checks if the data is at least min characters long
func MinLength(data string, min int) func(string) (string, bool) {
	return func(s string) (string, bool) {
		msg := fmt.Sprintf("%s must be at least %d characters long", s, min)
		if len(data) < min {
			return msg, false
		}
		return "", true
	}
}

// In checks if the data is in the collection
func In[T comparable](item T, collection ...T) func(string) (string, bool) {
	return func(s string) (string, bool) {
		for _, t := range collection {
			if item == t {
				return "", true
			}
		}
		return fmt.Sprintf("%v is not in any of the valid value", item), false
	}
}

// IsEmail checks if the data is a valid email address
func IsEmail(email string) func(string) (string, bool) {
	return func(s string) (string, bool) {
		_, err := mail.ParseAddress(email)
		return fmt.Sprintf("%s is not a valid email address", email), err == nil
	}
}

// IsAlphanumeric checks if the data is alphanumeric excluding space
func IsAlphanumeric(value string) func(string) (string, bool) {
	return func(s string) (string, bool) {
		if !regexp.MustCompile(`^([a-zA-Z0-9])+$`).Match([]byte(value)) {
			return fmt.Sprintf("%s must be alphanumeric", s), false
		}
		return "", true
	}
}
