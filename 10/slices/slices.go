package slices

import (
	"fmt"
	"reflect"
)

// ReturnSlice takes in any type and returns it if a slice or provides warning
func ReturnSlice(compLiteral interface{}) string {
	typ := reflect.ValueOf(compLiteral)
	switch {
	case typ.Kind() != reflect.Slice:
		return "This is not a slice"
	}
	return fmt.Sprintf("%v", compLiteral)
}
