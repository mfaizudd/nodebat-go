# NoDebat-go
![Coverage](https://img.shields.io/badge/Coverage-86.0%25-brightgreen)

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
    v.Builder("name", s.name).Required().IsAlphanumeric()

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

| Validator      | Tag             | Notes                                     |
| -------------- | --------------- | ----------------------------------------- |
| Required       | required        |                                           |
| IsAlphanumeric | is_alphanumeric |                                           |
| MinLength      | min_length      |                                           |
| MaxLength      | max_length      |                                           |
| Min            | min             | MinInt, MinUint, and MinFloat on builder  |
| Max            | max             | MinInt, MinUint, and MinFloat on builder  |
| Range          | range           |                                           |
| OneOf          | one_of          |                                           |
| IsEmail        | is_email        |                                           |
| IsISO8601      | is_iso8601      |                                           |
| IsISO8601Date  | is_iso8601_date |                                           |
| IsPhone        | is_phone        |                                           |
| IsUUID         | is_uuid         |                                           |
| MinDate        | min_date        | Attempts to parse string if using builder |
| MaxDate        | max_date        | Attempts to parse string if using builder |
| BetweenDate    | between_date    | Attempts to parse string if using builder |

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
v.Builder("field", someValue).Custom(SomeCustomValidator("arg1", someValue))
```

Or even create a validator directly

```go
v.Builder("field", someValue).Custom(func(field string) *validation.FieldError {
    return nil
})
// or, if not using builder
v.Add("field", func(field string) *validation.FieldError {
    return nil
})
```

## Why?
1. Fun,
2. It's more flexible than [package validator](https://github.com/go-playground/validator), I think.
With this library I can use whatever logic I want in my validation.
Like, for example, I can only validate a field if other field is empty, or
I can validate a field against a database by providing a database connection
dependency in the validation parameters, etc.
3. It's simple, my brain can understand this so you can too.
4. It still has `tags` to use with [universal-translator](https://github.com/go-playground/universal-translator).
I actually made this library because I'm not satisfied 
with [package validator](https://github.com/go-playground/validator)
and I only just want to validate some simple thing anyway.
5. I need a validation library that can return something like what laravel's validation returns.
