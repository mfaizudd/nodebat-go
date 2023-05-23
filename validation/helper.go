package validation

import (
	"strconv"
	"time"
)

// getTime parses a time.Time from a string or time.Time and returns the time.Time and a bool indicating if the parsing was successful
// If the parsing was not successful, the time.Time will be the zero value
// If the value is a string, it will attempt to parse it using the following layouts (in order):
// - time.RFC3339
// - "2006-01-02"
// - "2006-01-02 15:04:05"
// - "2006-01-02T15:04:05Z"
// - "2006-01-02T15:04:05"
// - "01/02/2006"
// - "01-02-2006"
// - "01/02/2006 15:04:05"
// - "01-02-2006 15:04:05"
func (v *Builder) getTime() (time.Time, bool) {
	switch val := v.value.(type) {
	case time.Time:
		return val, true
	case string, *string:
		stringval, ok := v.getString()
		if !ok {
			return time.Time{}, false
		}
		parsingLayouts := []string{
			time.RFC3339,
			"2006-01-02",
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05Z",
			"2006-01-02T15:04:05",
			"01/02/2006",
			"01-02-2006",
			"01/02/2006 15:04:05",
			"01-02-2006 15:04:05",
		}
		for _, layout := range parsingLayouts {
			if t, err := time.Parse(layout, stringval); err == nil {
				return t, true
			}
		}
	}
	return time.Time{}, false
}

func ptr[T any](v T) *T { return &v }

func ptrGet[T any](v *T, def T) (T, bool) {
	if v == nil {
		return def, false
	}
	return *v, true
}

func (v *Builder) hasError() bool {
	return v.validation.fieldErrors[v.field] != nil
}

func (v *Builder) getString() (string, bool) {
	var value string
	switch val := v.value.(type) {
	case string:
		value = val
	case *string:
		if val == nil {
			return "", false
		}
		value = *val
	default:
		return "", false
	}
	return value, true
}

func (v *Builder) getInt() (int64, bool) {
	var value int64
	switch val := v.value.(type) {
	case int:
		value = int64(val)
	case int8:
		value = int64(val)
	case int16:
		value = int64(val)
	case int32:
		value = int64(val)
	case int64:
		value = val
	case *int:
		ptrval, ok := ptrGet(val, 0)
		return int64(ptrval), ok
	case *int8:
		ptrval, ok := ptrGet(val, 0)
		return int64(ptrval), ok
	case *int16:
		ptrval, ok := ptrGet(val, 0)
		return int64(ptrval), ok
	case *int32:
		ptrval, ok := ptrGet(val, 0)
		return int64(ptrval), ok
	case *int64:
		ptrval, ok := ptrGet(val, 0)
		return int64(ptrval), ok
	case uint, uint8, uint16, uint32, uint64, *uint, *uint8, *uint16, *uint32, *uint64:
		uintval, ok := v.getUint()
		if !ok {
			return 0, false
		}
		value = int64(uintval)
	case float32, float64, *float32, *float64:
		floatval, ok := v.getFloat()
		if !ok {
			return 0, false
		}
		value = int64(floatval)
	case string, *string:
		stringval, ok := v.getString()
		if !ok {
			return 0, false
		}
		value, err := strconv.ParseInt(stringval, 10, 64)
		if err != nil {
			return 0, false
		}
		return value, true
	default:
		return 0, false
	}
	return value, true
}

func (v *Builder) getUint() (uint64, bool) {
	var value uint64
	switch val := v.value.(type) {
	case uint:
		value = uint64(val)
	case uint8:
		value = uint64(val)
	case uint16:
		value = uint64(val)
	case uint32:
		value = uint64(val)
	case uint64:
		value = val
	case *uint:
		ptrval, ok := ptrGet(val, 0)
		return uint64(ptrval), ok
	case *uint8:
		ptrval, ok := ptrGet(val, 0)
		return uint64(ptrval), ok
	case *uint16:
		ptrval, ok := ptrGet(val, 0)
		return uint64(ptrval), ok
	case *uint32:
		ptrval, ok := ptrGet(val, 0)
		return uint64(ptrval), ok
	case *uint64:
		ptrval, ok := ptrGet(val, 0)
		return ptrval, ok
	case int, int8, int16, int32, int64, *int, *int8, *int16, *int32, *int64:
		intval, ok := v.getInt()
		return uint64(intval), ok
	case float32, float64, *float32, *float64:
		floatval, ok := v.getFloat()
		return uint64(floatval), ok
	case string, *string:
		stringval, ok := v.getString()
		if !ok {
			return 0, false
		}
		value, err := strconv.ParseUint(stringval, 10, 64)
		if err != nil {
			return 0, false
		}
		return value, true
	default:
		return 0, false
	}
	return value, true
}

func (v *Builder) getFloat() (float64, bool) {
	var value float64
	switch val := v.value.(type) {
	case float32:
		value = float64(val)
	case float64:
		value = val
	case *float32:
		ptrval, ok := ptrGet(val, 0)
		return float64(ptrval), ok
	case *float64:
		ptrval, ok := ptrGet(val, 0)
		return ptrval, ok
	case int, int8, int16, int32, int64, *int, *int8, *int16, *int32, *int64:
		intval, ok := v.getInt()
		return float64(intval), ok
	case uint, uint8, uint16, uint32, uint64, *uint, *uint8, *uint16, *uint32, *uint64:
		uintval, ok := v.getUint()
		return float64(uintval), ok
	case string, *string:
		stringval, ok := v.getString()
		if !ok {
			return 0, false
		}
		value, err := strconv.ParseFloat(stringval, 64)
		if err != nil {
			return 0, false
		}
		return value, true
	default:
		return 0, false
	}
	return value, true
}
