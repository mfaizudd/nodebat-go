package validators

import (
	"github.com/mfaizudd/nodebat-go/validation"
	"testing"
)

func TestRequiredValidator(t *testing.T) {
	validValues := []string{
		"test",
		"test test",
		"test test test",
	}
	for _, value := range validValues {
		v := validation.New()
		v.Add("test", Required(value))
		if v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, value)
		}
	}
	v := validation.New()
	v.Add("test", Required(""))
	if v.Error() == nil {
		t.Fatalf(`value %q is valid, it should be invalid`, "")
	}
}

func TestIsAlphanumeric(t *testing.T) {
	validValues := []string{
		"test",
		"testtest",
		"testtesttest",
	}
	for _, value := range validValues {
		v := validation.New()
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
		v := validation.New()
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
		v := validation.New()
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
		v := validation.New()
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
		v := validation.New()
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
		v := validation.New()
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
		v := validation.New()
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
		v := validation.New()
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
		v := validation.New()
		v.Add("test", MaxLength(testCase.value, testCase.max))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid (max %d)`, testCase.value, testCase.max)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid (max %d)`, testCase.value, testCase.max)
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
		v := validation.New()
		v.Add("test", OneOf(testCase.value, testCase.allowed...))
		if testCase.expected && v.Error() != nil {
			t.Fatalf(`value %q is invalid, it should be valid`, testCase.value)
		}
		if !testCase.expected && v.Error() == nil {
			t.Fatalf(`value %q is valid, it should be invalid`, testCase.value)
		}
	}
}
