package validation

import "testing"

func TestValidation(t *testing.T) {
	v := New()
	v.Add("foo", func(s string) (string, bool) {
		return "foo is not valid", false
	})
	v.Add("bar", func(s string) (string, bool) {
		return "bar is not valid", false
	})
	if err := v.Error(); err == nil {
		t.Error("expected error, got nil")
	} else {
		if err.Error() != "foo is not valid, bar is not valid" {
			t.Errorf("expected error message 'foo is not valid, bar is not valid', got '%s'", err.Error())
		}
	}
}
