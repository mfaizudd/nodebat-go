package validation

import (
	"testing"
	"time"
)

func TestValidateIsCorrect(t *testing.T) {
	v := New()

	//         field, value
	v.Builder("minmax", 6).
		Min(5).
		Max(10)

	v.Builder("range", 6).
		Range(5, 10)

	v.Builder("required", "test").
		Required()

	v.Builder("minmaxlength", "test").
		MinLength(3).
		MaxLength(10)

	v.Builder("rangelength", "test").
		Length(3, 10)

	v.Builder("oneof", "test").
		OneOf("test", "test2")

	v.Builder("email", "email@domain.com").
		IsEmail()

	v.Builder("is_alphanumeric", "test123").
		IsAlphanumeric()

	v.Builder("is_iso8601", "2014-01-01T00:00:00Z").
		IsISO8601()

	v.Builder("is_iso8601date", "2014-01-01").
		IsISO8601Date()

	v.Builder("is_phone", "1234567890").
		IsPhone()

	v.Builder("is_uuid", "6ba7b810-9dad-11d1-80b4-00c04fd430c8").
		IsUUID()

	v.Builder("is_only_digits", "1234567890").
		IsOnlyDigits()

	v.Builder("min_max_date", time.Now()).
		MinDate(time.Now().AddDate(0, 0, -1)).
		MaxDate(time.Now().AddDate(0, 0, 1))

	v.Builder("between_date", time.Now()).
		BetweenDate(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 1))

    array := []string{"test", "test2"}

    v.Builder("min_max_count", array).
        MinCount(1).
        MaxCount(10)

	err := v.Error()
	if err != nil {
		t.Error("Expected error to be nil, got: ", err)
	}
}

func TestBuilder(t *testing.T) {
	v := New()
	vb := NewBuilder(v, "field", "")
	vb.Required()
	err := v.Error()
	if err == nil {
		t.Error("Expected error to be not nil")
	}
}

func TestValidateType(t *testing.T) {
	v := New()
	vb := NewBuilder(v, "field", "6")
	vb.Min(5)
	err := v.Error()
	if err == nil {
		t.Error("Expected error to be not nil")
	}
}
