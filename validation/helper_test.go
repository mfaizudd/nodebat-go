package validation

import "testing"

func TestParseTime(t *testing.T) {
	v := New()
	b := v.Builder("test", "test")
	// Test valid time
	if _, ok := parseTime("2006-01-02T15:04:05Z", b); !ok {
		t.Error("Expected valid time")
	}
	if _, ok := parseTime("2006-01-02", b); !ok {
		t.Error("Expected valid time")
	}

	// Test invalid time
	if _, ok := parseTime("2006-01-02T15:04:05", b); ok {
		t.Error("Expected invalid time")
	}
}
