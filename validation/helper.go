package validation

import (
	"fmt"
	"time"
)

// validateType validates the type of a value and adds an error to the validation object if it is not of the correct type
func validateType[T any](value interface{}, builder *Builder) (T, bool) {
	v, ok := value.(T)
	if !ok {
		builder.add(fmt.Sprintf("Invalid type: expected %T got %T, field: %v", v, value, builder.field), "invalid_type")
	}
	return v, ok
}

// parseTime parses a time.Time from a string or time.Time and returns the time.Time and a bool indicating if the parsing was successful
// If the parsing was not successful, the time.Time will be the zero value
// If the value is a string, it will attempt to parse it using the following layouts (in order):
// - time.RFC3339
// - "2006-01-02"
func parseTime(value interface{}, builder *Builder) (time.Time, bool) {
	switch val := value.(type) {
	case time.Time:
		return val, true
	case string:
		parsingLayouts := []string{
			time.RFC3339,
			"2006-01-02",
		}
		for _, layout := range parsingLayouts {
			if t, err := time.Parse(layout, val); err == nil {
				return t, true
			}
		}
	}
	builder.add(fmt.Sprintf("Invalid time: %v, field: %v", value, builder.field), "invalid_time")
	return time.Time{}, false
}
