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
- `Required`
- `IsAlphanumeric`
- `MinLength`
- `In`
- `IsEmail`

## Custom validators
To create a custom validation, you simply need to create a function that 
returns `func(f string) (m string, s bool)` where `f` is field name, `m`
is error message, and `s` is whether the checks is valid

Example: 
```go
// IsEmail checks if the data is a valid email address
func IsEmail(email string) func(string) (string, bool) {
	return func(s string) (string, bool) {
		_, err := mail.ParseAddress(email)
		return fmt.Sprintf("%s is not a valid email address", email), err == nil
	}
}
```
