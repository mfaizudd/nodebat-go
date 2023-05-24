package validation

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	v := New()
	validTimes := []interface{}{
		"2000-12-02T15:04:05Z",
		"2000-12-02",
		time.Now(),
		"2001-12-02 15:04:05",
		"2010-12-02T15:04:05",
		"12/02/2000",
		"12-02-2000",
		"12/02/2000 15:04:05",
		"12-02-2000 15:04:05",
	}
	// Test valid time
	for _, validTime := range validTimes {
		if val, ok := v.Builder("test", validTime).getTime(); !ok {
			t.Errorf("Expected valid time, got %v", val)
		}
	}

	// Test invalid time
	if val, ok := v.Builder("test", "lol").getTime(); ok {
		t.Errorf("Expected invalid time, got %v", val)
	}
}

func TestGetString(t *testing.T) {
	v := New()
	// Test valid string
	if _, ok := v.Builder("test", "test").getString(); !ok {
		t.Error("Expected valid string")
	}
	if _, ok := v.Builder("test", "").getString(); !ok {
		t.Error("Expected valid string")
	}

	type nilfield struct {field *string}
	test := &nilfield{}
	// Test invalid string
	if _, ok := v.Builder("test", test.field).getString(); ok {
		t.Error("Expected invalid string")
	}
}

var numbers = []interface{}{
	1,
	"1",
	int8(1),
	int16(1),
	int32(1),
	int64(1),
	uint(1),
	uint8(1),
	uint16(1),
	uint32(1),
	uint64(1),
	float32(1),
	float64(1),
	ptr(1),
	ptr("1"),
	ptr(uint(1)),
	ptr[int8](1),
	ptr[int16](1),
	ptr[int32](1),
	ptr[int64](1),
	ptr[uint8](1),
	ptr[uint16](1),
	ptr[uint32](1),
	ptr[uint64](1),
	ptr[float32](1),
	ptr[float64](1),
}

var nan = []interface{}{
	"test",
	"12-12-2012",
	"1.1.1",
	time.Now(),
	struct{}{},
}

func TestGetInt(t *testing.T) {
	v := New()

	// Test valid int
	for _, validInt := range numbers {
		if val, ok := v.Builder("test", validInt).getInt(); !ok {
			t.Errorf("Expected valid int, got %v", val)
		}
	}

	// Test invalid int
	for _, invalidInt := range nan {
		if val, ok := v.Builder("test", invalidInt).getInt(); ok {
			t.Errorf("Expected invalid int, got %v", val)
		}
	}
}

func TestGetUint(t *testing.T) {
	v := New()

	// Test valid uint
	for _, validUint := range numbers {
		if val, ok := v.Builder("test", validUint).getUint(); !ok {
			t.Errorf("Expected valid uint, got %v", val)
		}
	}

	// Test invalid uint
	for _, invalidUint := range nan {
		if val, ok := v.Builder("test", invalidUint).getUint(); ok {
			t.Errorf("Expected invalid uint, got %v", val)
		}
	}
}

func TestGetFloat(t *testing.T) {
	v := New()

	// Test valid float
	for _, validFloat := range numbers {
		if val, ok := v.Builder("test", validFloat).getFloat(); !ok {
			t.Errorf("Expected valid float, got %v", val)
		}
	}

	// Test invalid float
	for _, invalidFloat := range nan {
		if val, ok := v.Builder("test", invalidFloat).getFloat(); ok {
			t.Errorf("Expected invalid float, got %v", val)
		}
	}
}
