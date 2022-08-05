// Package pointers demonstrates pointers in go
package pointers

// GetAddressOfValue takes any value of any type and returns the memory address
//
// The address is accessed with:
//	&value
// And returned as the type pointer of any type with:
//	*interface{}
func GetAddressOfValue(value interface{}) *interface{} {
	return &value
}

// AssignToAddress takes a string and assigns a new pointer to the same address
func AssignToAddress(variable string) string {
	var result *string = &variable
	return result
}