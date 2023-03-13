package validation

import "testing"

func TestValidation(t *testing.T) {
	v := New()
	v.Add("foo", func(field string) *FieldError {
		return &FieldError{
			field:   field,
			message: "foo is not valid",
			tag:     "foo",
			value:   "foo",
		}
	})
	v.Add("bar", func(field string) *FieldError {
		return &FieldError{
			field:   field,
			message: "bar is not valid",
			tag:     "bar",
			value:   "foo",
		}
	})
	if err := v.Error(); err == nil {
		t.Error("expected error, got nil")
	} else if verr := err.(Error); len(verr.Errors()) != 2 {
        t.Errorf("expected 2 errors, got %d", len(verr.Errors()))
	}
}

func TestValidationReturnCurrectNumberOfErrors(t *testing.T) {
	v := New()
	v.Add("foo", func(field string) *FieldError {
		return &FieldError{
			field:   field,
			message: "foo is not valid first",
			tag:     "foo",
		}
	})
	v.Add("foo", func(field string) *FieldError {
		return &FieldError{
			field:   field,
			message: "foo is not valid again",
			tag:     "foo",
		}
	})
	v.Add("bar", func(field string) *FieldError {
		return &FieldError{
			field:   field,
			message: "bar is not valid",
			tag:     "bar",
		}
	})
	if err := v.Error(); err == nil {
		t.Error("expected error, got nil")
		// if err is not nil, then it should be of type validation Error
	} else if verr := err.(Error); len(verr.Errors()) != 2 {
		t.Errorf("expected 2 errors, got %d", len(verr.Errors()))
	}
}
