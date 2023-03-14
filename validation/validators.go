package validation

import (
	"fmt"
	"net/mail"
	"reflect"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Validator func(field string) *FieldError

// Required checks if the data is nil or empty string
func Required(data interface{}) Validator {
	return func(field string) *FieldError {
		msg := fmt.Sprintf("%s is required", field)
		if data == nil || (reflect.ValueOf(data).Kind() == reflect.Ptr && reflect.ValueOf(data).IsNil()) {
			return NewFieldError(field, msg, "required", nil)
		}
		if str, ok := data.(string); ok && str == "" {
			return NewFieldError(field, msg, "required", str)
		}
		return nil
	}
}

// Min checks if the data is at least min
//
// Has one parameter: min (any number type, except complex)
func Min[T Number](data T, min T) Validator {
	return func(field string) *FieldError {
		msg := fmt.Sprintf("%s must be at least %v", field, min)
		if data < min {
			err := NewFieldError(field, msg, "min", data)
			err.SetParam("min", min)
			return err
		}
		return nil
	}
}

// Max checks if the data is at most max
//
// Has one parameter: max (any number type, except complex)
func Max[T Number](data T, max T) Validator {
	return func(field string) *FieldError {
		msg := fmt.Sprintf("%s must be at most %v", field, max)
		if data > max {
			err := NewFieldError(field, msg, "max", data)
			err.SetParam("max", max)
			return err
		}
		return nil
	}
}

// Range checks if the data is between min and max
//
// Has two parameters: min (any number type, except complex), max (same type as min)
func Range[T Number](data T, min T, max T) Validator {
	return func(field string) *FieldError {
		msg := fmt.Sprintf("%s must be between %v and %v", field, min, max)
		if data < min || data > max {
			err := NewFieldError(field, msg, "range", data)
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
func MinLength(data string, min int) Validator {
	return func(field string) *FieldError {
		msg := fmt.Sprintf("%s must be at least %d characters long", field, min)
		if len(data) < min {
			err := NewFieldError(field, msg, "min_length", data)
			err.SetParam("min", min)
			return err
		}
		return nil
	}
}

// MaxLength checks if the data is at most max characters long
//
// Has one parameter: max (int)
func MaxLength(data string, max int) Validator {
	return func(field string) *FieldError {
		msg := fmt.Sprintf("%s must be at most %d characters long", field, max)
		if len(data) > max {
			err := NewFieldError(field, msg, "max_length", data)
			err.SetParam("max", max)
			return err
		}
		return nil
	}
}

// Length checks if the data is between min and max characters long
//
// Has one parameter: length (int)
func Length(data string, min int, max int) Validator {
	return func(field string) *FieldError {
		msg := fmt.Sprintf("%s must be between %d and %d characters long", field, min, max)
		if len(data) < min || len(data) > max {
			err := NewFieldError(field, msg, "length", data)
			err.SetParam("min", min)
			err.SetParam("max", max)
			return err
		}
		return nil
	}
}

// OneOf checks if the data is in the collection
//
// Has one parameter named "collection" which is a slice of strings
func OneOf(item string, collection ...string) Validator {
	return func(field string) *FieldError {
		for _, it := range collection {
			if item == it {
				return nil
			}
		}
		msg := fmt.Sprintf("%s is not in the collection", field)
		err := NewFieldError(field, msg, "one_of", item)
		err.SetParam("collection", collection)
		return err
	}
}

// IsEmail checks if the data is a valid email address
func IsEmail(email string) Validator {
	return func(field string) *FieldError {
		_, emailErr := mail.ParseAddress(email)
		if emailErr != nil {
			msg := fmt.Sprintf("%s is not a valid email address", field)
			return NewFieldError(field, msg, "is_email", email)
		}
		return nil
	}
}

// IsAlphanumeric checks if the data is alphanumeric excluding space
func IsAlphanumeric(value string) Validator {
	return func(field string) *FieldError {
		if !regexp.MustCompile(`^([a-zA-Z0-9])+$`).Match([]byte(value)) {
			msg := fmt.Sprintf("%s must be alphanumeric", field)
			return NewFieldError(field, msg, "is_alphanumeric", value)
		}
		return nil
	}
}

// IsISO8601 checks if the data is a valid ISO8601 date
func IsISO8601(date string) Validator {
	return func(field string) *FieldError {
		if _, err := time.Parse(time.RFC3339, date); err != nil {
			msg := fmt.Sprintf("%s is not a valid ISO8601 date", field)
			return NewFieldError(field, msg, "is_iso8601", date)
		}
		return nil
	}
}

// IsISO8601Date checks if the data is a valid ISO8601 date
func IsISO8601Date(date string) Validator {
	return func(field string) *FieldError {
		if _, err := time.Parse("2006-01-02", date); err != nil {
			msg := fmt.Sprintf("%s is not a valid ISO8601 date", field)
			return NewFieldError(field, msg, "is_iso8601_date", date)
		}
		return nil
	}
}

// IsPhone checks if the data is a valid phone number
func IsPhone(phone string) Validator {
	return func(field string) *FieldError {
		if !regexp.MustCompile(`^(\+?)([0-9])+$`).Match([]byte(phone)) {
			msg := fmt.Sprintf("%s is not a valid phone number", field)
			return NewFieldError(field, msg, "is_phone", phone)
		}
		return nil
	}
}

// IsUUID checks if the data is a valid UUID
func IsUUID(input string) Validator {
	return func(field string) *FieldError {
		if _, err := uuid.Parse(input); err != nil {
			msg := fmt.Sprintf("%s is not a valid UUID", field)
			return NewFieldError(field, msg, "is_uuid", input)
		}
		return nil
	}
}

// IsOnlyDigits checks if the data contains only digits
func IsOnlyDigits(input string) Validator {
	return func(field string) *FieldError {
		if !regexp.MustCompile(`^[0-9]+$`).Match([]byte(input)) {
			msg := fmt.Sprintf("%s contains non-digit characters", field)
			return NewFieldError(field, msg, "is_only_digits", input)
		}
		return nil
	}
}

// MinDate checks if the date is after the given date
//
// Has one parameter: minDate (time.Time)
func MinDate(date time.Time, minDate time.Time) Validator {
	return func(field string) *FieldError {
		if date.Before(minDate) {
			msg := fmt.Sprintf("%s is before %s", field, minDate)
			err := NewFieldError(field, msg, "min_date", date)
			err.SetParam("min_date", minDate)
			return err
		}
		return nil
	}
}

// MaxDate checks if the date is before the given date
//
// Has one parameter: maxDate (time.Time)
func MaxDate(date time.Time, maxDate time.Time) Validator {
	return func(field string) *FieldError {
		if date.After(maxDate) {
			msg := fmt.Sprintf("%s is after %s", field, maxDate)
			err := NewFieldError(field, msg, "max_date", date)
			err.SetParam("max_date", maxDate)
			return err
		}
		return nil
	}
}

// BetweenDate checks if the date is between the given dates
//
// Has two parameters: minDate (time.Time), maxDate (time.Time)
func BetweenDate(date time.Time, minDate time.Time, maxDate time.Time) Validator {
	return func(field string) *FieldError {
		if date.Before(minDate) || date.After(maxDate) {
			msg := fmt.Sprintf("%s is not between %s and %s", field, minDate, maxDate)
			err := NewFieldError(field, msg, "between_date", date)
			err.SetParam("min_date", minDate)
			err.SetParam("max_date", maxDate)
			return err
		}
		return nil
	}
}

// MinCount checks if the length of the array/slice/map is greater than or equal to the given number
//
// Has one parameter: min (int)
func MinCount(array interface{}, min int) Validator {
	return func(field string) *FieldError {
		switch reflect.TypeOf(array).Kind() {
		case reflect.Slice, reflect.Array, reflect.Map:
			v := reflect.ValueOf(array)
			if v.Len() < min {
				msg := fmt.Sprintf("%s must have at least %d items", field, min)
				err := NewFieldError(field, msg, "min_count", array)
				err.SetParam("min_count", min)
				return err
			}
		default:
			msg := fmt.Sprintf("%s must be an array or slice", field)
			err := NewFieldError(field, msg, "min_count", array)
			err.SetParam("min_count", min)
			return err
		}
		return nil
	}
}

// MaxCount checks if the length of the array/slice/map is less than or equal to the given number
//
// Has one parameter: max (int)
func MaxCount(array interface{}, max int) Validator {
	return func(field string) *FieldError {
		switch reflect.TypeOf(array).Kind() {
		case reflect.Slice, reflect.Array, reflect.Map:
			v := reflect.ValueOf(array)
			if v.Len() > max {
				msg := fmt.Sprintf("%s must have at most %d items", field, max)
				err := NewFieldError(field, msg, "max_count", array)
				err.SetParam("max_count", max)
				return err
			}
		default:
			msg := fmt.Sprintf("%s must be an array or slice", field)
			err := NewFieldError(field, msg, "max_count", array)
			err.SetParam("max_count", max)
			return err
		}
		return nil
	}
}
