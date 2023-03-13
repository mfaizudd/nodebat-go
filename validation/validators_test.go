package validation

import (
	"testing"
	"time"
)

func TestRequiredValidator(t *testing.T) {
	validValues := []string{
		"test",
		"test test",
		"test test test",
	}
	for _, value := range validValues {
		v := New()
		v.Add("test", Required(value))
		if v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, value)
		}
	}
	v := New()
	v.Add("test", Required(""))
	if v.Error() == nil {
		t.Fatalf(`value %q is valid, it should be invalid`, "")
	}
}

func TestRequiredInterface(t *testing.T) {
	someStruct := struct{}{}
	testCases := []struct {
		value interface{}
		valid bool
	}{
		{nil, false},
		{interface{}(nil), false},
		{(*string)(nil), false},
		{"", false},
		{"test", true},
		{someStruct, true},
		{(*struct{})(nil), false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", Required(testCase.value))
		if testCase.valid && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.valid && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestIsAlphanumeric(t *testing.T) {
	validValues := []string{
		"test",
		"testtest",
		"testtesttest",
	}
	for _, value := range validValues {
		v := New()
		v.Add("test", IsAlphanumeric(value))
		if v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, value)
		}
	}
	invalidValues := []string{
		"te st",
		"test@",
		"test#",
		"test$",
		"test%",
		"test^",
		"test&",
		"test*",
		"test(",
		"test)",
		"test-",
		"test_",
		"test+",
		"test=",
		"test{",
		"test}",
		"test[",
		"test]",
		"test|",
		"test\\",
		"test:",
		"test;",
		"test\"",
		"test'",
		"test<",
		"test>",
		"test,",
		"test.",
		"test?",
		"test/",
	}
	for _, value := range invalidValues {
		v := New()
		v.Add("test", IsAlphanumeric(value))
		if v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, value)
		}
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		value    int
		min      int
		expected bool
	}{
		{1, 1, true},
		{1, 2, false},
		{2, 1, true},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", Min(testCase.value, testCase.min))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %d is invalid, it should be valid (min %d)`, testCase.value, testCase.min)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %d is valid, it should be invalid (min %d)`, testCase.value, testCase.min)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		value    int
		max      int
		expected bool
	}{
		{1, 1, true},
		{1, 2, true},
		{2, 1, false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", Max(testCase.value, testCase.max))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %d is invalid, it should be valid (max %d)`, testCase.value, testCase.max)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %d is valid, it should be invalid (max %d)`, testCase.value, testCase.max)
		}
	}
}

func TestRange(t *testing.T) {
	testCases := []struct {
		value    int
		min      int
		max      int
		expected bool
	}{
		{1, 1, 1, true},
		{1, 1, 2, true},
		{1, 2, 3, false},
		{2, 1, 3, true},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", Range(testCase.value, testCase.min, testCase.max))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %d is invalid, it should be valid (range %d - %d)`, testCase.value, testCase.min, testCase.max)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %d is valid, it should be invalid (range %d - %d)`, testCase.value, testCase.min, testCase.max)
		}
	}
}

func TestIsEmail(t *testing.T) {
	validValues := []string{
		"test@email.com",
		"another@email",
		"go@goo.goo.go",
	}
	for _, value := range validValues {
		v := New()
		v.Add("test", IsEmail(value))
		if v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, value)
		}
	}
	invalidValues := []string{
		"testemail.com",
		"testemail",
		"testemail@",
	}
	for _, value := range invalidValues {
		v := New()
		v.Add("test", IsEmail(value))
		if v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, value)
		}
	}
}

func TestMinLength(t *testing.T) {
	testCases := []struct {
		value    string
		min      int
		expected bool
	}{
		{"test", 3, true},
		{"test", 2, true},
		{"test", 1, true},
		{"test", 0, true},
		{"", 0, true},
		{"", 1, false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", MinLength(testCase.value, testCase.min))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestMaxLength(t *testing.T) {
	testCases := []struct {
		value    string
		max      int
		expected bool
	}{
		{"test", 4, true},
		{"test", 3, false},
		{"test", 2, false},
		{"test", 1, false},
		{"test", 0, false},
		{"", 0, true},
		{"", 1, true},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", MaxLength(testCase.value, testCase.max))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid (max %d)`, testCase.value, testCase.max)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid (max %d)`, testCase.value, testCase.max)
		}
	}
}

func TestLength(t *testing.T) {
	testCases := []struct {
		value    string
		min      int
		max      int
		expected bool
	}{
		{"test", 4, 4, true},
		{"test", 3, 4, true},
		{"test", 2, 3, false},
		{"test", 5, 6, false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", Length(testCase.value, testCase.min, testCase.max))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid (min %d, max %d)`, testCase.value, testCase.min, testCase.max)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid (min %d, max %d)`, testCase.value, testCase.min, testCase.max)
		}
	}
}

func TestIn(t *testing.T) {
	testCases := []struct {
		value    string
		allowed  []string
		expected bool
	}{
		{"test", []string{"test"}, true},
		{"test", []string{"test", "test2"}, true},
		{"test", []string{"test2"}, false},
		{"test", []string{}, false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", OneOf(testCase.value, testCase.allowed...))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestIsISO8601(t *testing.T) {
	testCases := []struct {
		value    string
		expected bool
	}{
		{"2018-01-01T00:00:00Z", true},
		{"2018-01-01T00:00:00+00:00", true},
		{"2018-01-01T00:00:00+01:00", true},
		{"2018-01-01T00:00:00-01:00", true},
		{"testdate", false},
		{"2018-01-01T00:00:00+01", false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", IsISO8601(testCase.value))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestIsISO8601Date(t *testing.T) {
	testCases := []struct {
		value    string
		expected bool
	}{
		{"2018-01-01", true},
		{"testdate", false},
		{"2018-01-01T00:00:00+01:00", false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", IsISO8601Date(testCase.value))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestIsPhone(t *testing.T) {
	testCases := []struct {
		value    string
		expected bool
	}{
		{"+1 1234567890", false},
		{"+6281234567890", true},
		{"a phone", false},
		{"085211111111", true},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", IsPhone(testCase.value))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestIsUUID(t *testing.T) {
	testCases := []struct {
		value    string
		expected bool
	}{
		{"a uuid", false},
		{"6ba7b810-9dad-11d1-80b4-00c04fd430c8", true},
		{"aVeryLongStringThatIsNotAUUID", false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", IsUUID(testCase.value))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestIsOnlyDigits(t *testing.T) {
	testCases := []struct {
		value    string
		expected bool
	}{
		{"a string", false},
		{"123", true},
		{"-123", false}, // negative sign is not a digit
		{"aVeryLongStringThatIsNotOnlyDigits", false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", IsOnlyDigits(testCase.value))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestMinDate(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		value    time.Time
		expected bool
	}{
		{now, true},
		{now.AddDate(0, 0, -1), false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", MinDate(testCase.value, now))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestMaxDate(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		value    time.Time
		expected bool
	}{
		{now, true},
		{now.AddDate(0, 0, 1), false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", MaxDate(testCase.value, now))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestBetweenDate(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		value    time.Time
		min      time.Time
		max      time.Time
		expected bool
	}{
		{now, now.AddDate(0, 0, -1), now.AddDate(0, 0, 1), true},
		{now, now.AddDate(0, 0, 1), now.AddDate(0, 0, 2), false},
		{now, now.AddDate(0, 0, -2), now.AddDate(0, 0, -1), false},
	}
	for _, testCase := range testCases {
		v := New()
		v.Add("test", BetweenDate(testCase.value, testCase.min, testCase.max))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}

func TestMinCount(t *testing.T) {
    testCases := []struct {
        value []interface{}
        min int
        expected bool
    }{
        {[]interface{}{1, 2, 3}, 2, true},
        {[]interface{}{1, 2, 3}, 4, false},
    }
    for _, testCase := range testCases {
        v := New()
        v.Add("test", MinCount(testCase.value, testCase.min))
        if testCase.expected && v.Error() != nil {
            t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
        }
        if !testCase.expected && v.Error() == nil {
            t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
        }
    }
}

func TestMaxCount(t *testing.T) {
    testCases := []struct {
        value []interface{}
        max int
        expected bool
    }{
        {[]interface{}{1, 2, 3}, 4, true},
        {[]interface{}{1, 2, 3}, 2, false},
    }
    for _, testCase := range testCases {
        v := New()
        v.Add("test", MaxCount(testCase.value, testCase.max))
        if testCase.expected && v.Error() != nil {
            t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
        }
        if !testCase.expected && v.Error() == nil {
            t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
        }
    }
}
