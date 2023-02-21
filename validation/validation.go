package validation

type Validation struct {
	error    error
	fieldErrors map[string]*FieldError
}

func New() *Validation {
	errors := make(map[string]*FieldError, 0)
	return &Validation{error: nil, fieldErrors: errors}
}

func (v *Validation) Add(field string, validations ...func(field string) (err *FieldError, ok bool)) {
	for _, validation := range validations {
		if err, ok := validation(field); !ok && v.fieldErrors[field] == nil {
            v.fieldErrors[field] = err
		}
	}
}

func (v *Validation) Error() error {
	if len(v.fieldErrors) > 0 {
		v.error = NewError(v.fieldErrors)
	}
	return v.error
}