package validation

import "testing"

func TestValidation(t *testing.T) {
	v := New()
	v.Add("foo", func(field string) (err *FieldError, ok bool) {
		return &FieldError{
            field: field,
            message: "foo is not valid",
            tag: "foo",
            value: "foo",
        }, false
	})
	v.Add("bar", func(field string) (err *FieldError, ok bool) {
        return &FieldError{
            field: field,
            message: "bar is not valid",
            tag: "bar",
            value: "foo",
        }, false
    })
    if err := v.Error(); err == nil {
		t.Error("expected error, got nil")
	} else {
        expected := "foo: foo is not valid, bar: bar is not valid" 
        if err.Error() != expected {
			t.Errorf("expected error message '%s', got '%s'", expected, err.Error())
		}
	}
}

func TestValidationReturnCurrectNumberOfErrors(t *testing.T) {
	v := New()
	v.Add("foo", func(field string) (err *FieldError, ok bool) {
		return &FieldError{
            field: field,
            message: "foo is not valid first",
            tag: "foo",
        }, false
	})
	v.Add("foo", func(field string) (err *FieldError, ok bool) {
		return &FieldError{
            field: field,
            message: "foo is not valid again",
            tag: "foo",
        }, false
	})
	v.Add("bar", func(field string) (err *FieldError, ok bool) {
        return &FieldError{
            field: field,
            message: "bar is not valid",
            tag: "bar",
        }, false
    })
    if err := v.Error(); err == nil {
		t.Error("expected error, got nil")
	} else {
        expected := "foo: foo is not valid first, bar: bar is not valid"
        if err.Error() != expected {
            t.Errorf("expected error message '%s', got '%s'", expected, err.Error())
		}
	}
}