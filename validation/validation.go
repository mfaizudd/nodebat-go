package validation

type Validation struct {
	error    error
	messages map[string][]string
}

func New() *Validation {
	messages := make(map[string][]string, 0)
	return &Validation{error: nil, messages: messages}
}

func (v *Validation) Add(field string, validations ...func(string) (string, bool)) {
	for _, validation := range validations {
		if msg, ok := validation(field); !ok {
			v.messages[field] = append(v.messages[field], msg)
		}
	}
}

func (v *Validation) Error() error {
	if len(v.messages) > 0 {
		v.error = NewError(v.messages)
	}
	return v.error
}
