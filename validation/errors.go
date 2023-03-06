package validation

import (
	"fmt"
	"strings"
)

type Error struct {
	errors map[string]*FieldError
}

type FieldError struct {
	field   string
	message string
	tag     string
	value   interface{}
	params  map[string]interface{}
}

func NewFieldError(field, message, tag string, value interface{}) *FieldError {
	params := make(map[string]interface{})
	return &FieldError{field, message, tag, value, params}
}

func (e *FieldError) Field() string {
	return e.field
}

func (e *FieldError) Message() string {
	return e.message
}

func (e *FieldError) SetMessage(message string) {
	e.message = message
}

func (e *FieldError) SetParam(key string, value interface{}) {
	e.params[key] = value
}

func (e *FieldError) Param(key string) interface{} {
	return e.params[key]
}

func (e *FieldError) Params() map[string]interface{} {
	return e.params
}

func (e *FieldError) HasParams() bool {
	return len(e.params) > 0
}

func (e *FieldError) Tag() string {
	return e.tag
}

func (e *FieldError) Value() interface{} {
	return e.value
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.field, e.message)
}

func NewError(fieldErrors map[string]*FieldError) Error {
	return Error{fieldErrors}
}

func (e Error) Errors() map[string]*FieldError {
	if len(e.errors) > 0 {
		return e.errors
	}
	return nil
}

func (e Error) Error() string {
	if len(e.errors) <= 0 {
		return ""
	}
	messages := make([]string, 0)
	for _, err := range e.errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, ", ")
}
