# NoDebat-go
![Coverage](https://img.shields.io/badge/Coverage-93.9%25-brightgreen)

This is a really simple implementation of validation in go

## Basic usage

To validate a struct, you can just create a validation method like so:
```go
import "github.com/mfaizudd/nodebat-go/validation"

type Student struct {
    name string
}

func (s *Student) Validate() error {
    v := validation.New()
    v.Add("name", validation.Required(s.name), validation.IsAlphanumeric(s.name))
    return v.Error()
}
```

## Basic usage using builder

The builder struct provides a way to chain validators easier.
The above example can be recreated using builder like so:
```go
import "github.com/mfaizudd/nodebat-go/validation"

type Student struct {
    name string
}

func (s *Student) Validate() error {
    v := validation.New()

    //        field , value
    v.Builder("name", s.name).
        Required().
        IsAlphanumeric()

    return v.Error()
}
```

However, when using builder, you will lose the ability of
compile time checks of type in validators. for validators
like min, max, etc. where if the value is not the specific
type that's required by the validator, the validation will
fail and return `invalid_type` tag.

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
| Range          | range           |
| OneOf          | one_of          |
| IsEmail        | is_email        |
| IsISO8601      | is_iso8601      |
| IsISO8601Date  | is_iso8601_date |
| IsPhone        | is_phone        |
| IsUUID         | is_uuid         |
| MinDate        | min_date        |
| MaxDate        | max_date        |
| BetweenDate    | between_date    |

## Custom validators
To create a custom validation, you simply need to create a function that 
returns `func(field string) *validation.FieldError`.

`validation.Validator` is just a wrapper type of 
`func(field string) *validation.FieldError`.

Example: 
```go
// IsEmail checks if the data is a valid email address
func IsEmail(email string) validation.Validator {
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

## Validate using custom validators and builder
When you're using builder, you can add arbitrary validator
using the `Custom` method

```go
v.Builder().Custom(SomeCustomValidator("arg1"))
```
