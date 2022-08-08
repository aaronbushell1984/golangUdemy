// Package typeassertion demonstrates the use of type assertions in go
package typeassertion

// IsString returns true if argument is of type string
func IsString(s interface{}) bool {
	s, ok := s.(string)
	return ok
}

// IsIntOrInt64 switches on type to return true if argument is of type int or int64
func IsIntOrInt64(n interface{}) bool {
	switch n.(type) {
	case int, int64:
		return true
	default:
		return false
	}	
}