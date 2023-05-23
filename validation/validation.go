package validation

type Validation struct {
	error       error
	fieldErrors map[string]*FieldError
}

func New() *Validation {
	errors := make(map[string]*FieldError, 0)
	return &Validation{error: nil, fieldErrors: errors}
}

func (v *Validation) Builder(field string, value interface{}) *Builder {
	return NewBuilder(v, field, value)
}

func (v *Validation) Add(field string, validations ...Validator) {
	for _, validation := range validations {
		if v.fieldErrors[field] != nil {
			return
		}
		if err := validation(field); err != nil {
			v.fieldErrors[field] = err
		}
	}
}

func (v *Validation) AddError(field string, err *FieldError) {
	if v.fieldErrors[field] == nil {
		v.fieldErrors[field] = err
	}
}

func (v *Validation) Error() error {
	if len(v.fieldErrors) > 0 {
		v.error = NewError(v.fieldErrors)
	}
	return v.error
}
