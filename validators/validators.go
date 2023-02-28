package validators

import (
	"fmt"
	"net/mail"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/mfaizudd/nodebat-go/validation"
)

// Required checks if the data is nil or empty string
func Required(data interface{}) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		msg := fmt.Sprintf("%s is required", field)
		if data == nil {
			return validation.NewFieldError(field, msg, "required", nil)
		}
		if str, ok := data.(string); ok && str == "" {
			return validation.NewFieldError(field, msg, "required", str)
		}
		return nil
	}
}

// Min checks if the data is at least min
//
// Has one parameter: min (int)
func Min(data int, min int) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		msg := fmt.Sprintf("%s must be at least %d", field, min)
		if data < min {
			err := validation.NewFieldError(field, msg, "min", data)
			err.SetParam("min", min)
			return err
		}
		return nil
	}
}

// Max checks if the data is at most max
//
// Has one parameter: max (int)
func Max(data int, max int) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		msg := fmt.Sprintf("%s must be at most %d", field, max)
		if data > max {
			err := validation.NewFieldError(field, msg, "max", data)
			err.SetParam("max", max)
			return err
		}
		return nil
	}
}

// Range checks if the data is between min and max
//
// Has two parameters: min (int) and max (int)
func Range(data int, min int, max int) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		msg := fmt.Sprintf("%s must be between %d and %d", field, min, max)
		if data < min || data > max {
			err := validation.NewFieldError(field, msg, "range", data)
			err.SetParam("min", min)
			err.SetParam("max", max)
			return err
		}
		return nil
	}
}

// MinLength checks if the data is at least min characters long
//
// Has one parameter: min (int)
func MinLength(data string, min int) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		msg := fmt.Sprintf("%s must be at least %d characters long", field, min)
		if len(data) < min {
			err := validation.NewFieldError(field, msg, "min_length", data)
			err.SetParam("min", min)
			return err
		}
		return nil
	}
}

// MaxLength checks if the data is at most max characters long
//
// Has one parameter: max (int)
func MaxLength(data string, max int) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		msg := fmt.Sprintf("%s must be at most %d characters long", field, max)
		if len(data) > max {
			err := validation.NewFieldError(field, msg, "max_length", data)
			err.SetParam("max", max)
			return err
		}
		return nil
	}
}

// OneOf checks if the data is in the collection
//
// Has one parameter named "collection" which is a slice of strings
func OneOf[T comparable](item T, collection ...T) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		for _, it := range collection {
			if item == it {
				return nil
			}
		}
		msg := fmt.Sprintf("%s is not in the collection", field)
		err := validation.NewFieldError(field, msg, "one_of", item)
		err.SetParam("collection", collection)
		return err
	}
}

// IsEmail checks if the data is a valid email address
func IsEmail(email string) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		_, emailErr := mail.ParseAddress(email)
		if emailErr != nil {
			msg := fmt.Sprintf("%s is not a valid email address", field)
			return validation.NewFieldError(field, msg, "is_email", email)
		}
		return nil
	}
}

// IsAlphanumeric checks if the data is alphanumeric excluding space
func IsAlphanumeric(value string) func(field string) *validation.FieldError {
	return func(field string) *validation.FieldError {
		if !regexp.MustCompile(`^([a-zA-Z0-9])+$`).Match([]byte(value)) {
			msg := fmt.Sprintf("%s must be alphanumeric", field)
			return validation.NewFieldError(field, msg, "is_alphanumeric", value)
		}
		return nil
	}
}

// IsISO8601 checks if the data is a valid ISO8601 date
func IsISO8601(date string) func(field string) *validation.FieldError {
    return func(field string) *validation.FieldError {
        if _, err := time.Parse(time.RFC3339, date); err != nil {
            msg := fmt.Sprintf("%s is not a valid ISO8601 date", field)
            return validation.NewFieldError(field, msg, "is_iso8601", date)
        }
        return nil
    }
}

// IsISO8601Date checks if the data is a valid ISO8601 date
func IsISO8601Date(date string) func(field string) *validation.FieldError {
    return func(field string) *validation.FieldError {
        if _, err := time.Parse("2006-01-02", date); err != nil {
            msg := fmt.Sprintf("%s is not a valid ISO8601 date", field)
            return validation.NewFieldError(field, msg, "is_iso8601_date", date)
        }
        return nil
    }
}

// IsPhone checks if the data is a valid phone number
func IsPhone(phone string) func(field string) *validation.FieldError {
    return func(field string) *validation.FieldError {
        if !regexp.MustCompile(`^(\+?)([0-9])+$`).Match([]byte(phone)) {
            msg := fmt.Sprintf("%s is not a valid phone number", field)
            return validation.NewFieldError(field, msg, "is_phone", phone)
        }
        return nil
    }
}

// IsUUID checks if the data is a valid UUID
func IsUUID(input string) func(field string) *validation.FieldError {
    return func(field string) *validation.FieldError {
        if _, err := uuid.Parse(input); err != nil {
            msg := fmt.Sprintf("%s is not a valid UUID", field)
            return validation.NewFieldError(field, msg, "is_uuid", input)
        }
        return nil
    }
}
