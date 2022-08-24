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
	return *result
}

// GetDereference takes a string address which is of type pointer:
//	func Dereference(address *string) string
// The address is dereferenced using the opertaor *:
//	return *address
// NB. * in a type is:
//	pointer - (*string)
// * in dereferencing an address is:
//	operator - (*address)
func GetDereference(address *string) string {
	return *address
}

// ChangeValueViaNewIdentifier takes a value and uses a new identifier to change the value
//
// Assigning a new identifier to the address of the value:
//	newIdentifier := &value
// Dereference the new identifier to access the value and change:
//	*newIdentifier = "new"
func ChangeValueViaNewIdentifier(value string) string {
	newIdentifier := &value
	*newIdentifier = "new"
	return value
}

// GetChanged takes in a string and makes it "changed"
//
// Example shows that:
//	fmt.Println(GetChanged(Same))
// Has no affect on the value of Same after GetChanged is called because of function scope
func GetChanged(change string) string {
	change = "Changed"
	return change
}

// GetPointerChanged takes in an address of type pointer:
//	GetPointerChanged(change *string)
// Dereferencing the value and assigning it to a new value directly changes value stored in memory:
//	*change = "Changed"
// Note that the value is dereferenced again to meet return as string
//	return *change
func GetPointerChanged(change *string) string {
	*change = "Changed"
	return *change
}
