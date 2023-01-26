package validation

import (
	"fmt"
	"strings"
)

type Error struct {
	messages map[string][]string
}

func NewError(messages map[string][]string) Error {
	return Error{messages}
}

func (e Error) Errors() map[string][]string {
	if len(e.messages) > 0 {
		return e.messages
	}
	return nil
}

func (e Error) Error() string {
	if len(e.messages) <= 0 {
		return ""
	}
	var b strings.Builder
	for field, errs := range e.messages {
		_, err := fmt.Fprintf(&b, "%s:%s", field, strings.Join(errs, ","))
		if err != nil {
			return ""
		}
	}
	return b.String()
}
