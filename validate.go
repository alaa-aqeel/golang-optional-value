package optional

import (
	"reflect"
	"strings"
)

// IsEmpty checks if a value is "empty":
// - string → len == 0
// - slice → len == 0
// - map   → len == 0
// - array → len == 0
func IsEmpty[T any](v T) bool {
	val := reflect.ValueOf(v)

	switch val.Kind() {

	case reflect.String:
		// Trim spaces before checking length
		s := strings.TrimSpace(val.String())
		return len(s) == 0

	case reflect.Slice, reflect.Map, reflect.Array:
		return val.Len() == 0
	}

	return false
}

// IsZero checks if v is the zero-value of its type:
// - int → 0
// - float → 0
// - bool → false
// - struct → zero struct
func IsZero[T any](v T) bool {
	return reflect.ValueOf(v).IsZero()
}
