# NoDebat-go

This is a really simple implementation of validation in go

## Basic usage

To validate a struct, you can just create a validation method like so:
```go
import "github.com/mfaizudd/nodebat-go/validation"
import "github.com/mfaizudd/nodebat-go/validators"

type Student struct {
    name string
}

func (s *Student) Validate() error {
    v := validation.New()
    v.Add("name", validators.Required(s.name), validators.IsAlphanumeric(s.name))
    return v.Error()
}
```

## Built in validators
You can use tags to translate the error message

| Validator      | Tag             |
|----------------|-----------------|
| Required       | required        |
| IsAlphanumeric | is_alphanumeric |
| MinLength      | min_length      |
| MaxLength      | max_length      |
| Min            | min             |
| Max            | max             |
| OneOf          | one_of          |
| IsEmail        | is_email        |

## Custom validators
To create a custom validation, you simply need to create a function that 
returns `func(field string) *validation.FieldError` where `f` is field name, `m`
is error message, and `s` is whether the checks is valid

Example: 
```go
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
```
