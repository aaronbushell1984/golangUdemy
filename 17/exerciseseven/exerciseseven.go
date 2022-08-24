// Package exerciseseven is the seventh set of exercises in course
package exerciseseven

import "fmt"

type person struct {
	name string
}

// ReturnAddressType returns the type of the address of the supplied string
func ReturnAddressType(variable string) string {
	return fmt.Sprintf("%T", &variable)
}

// ChangeMe takes the address of a struct person of type pointer to person
//	ChangeMe(person *person)
// The name field is change to Mary using:
//	(*person).name = "Mary"
func ChangeMe(person *person) {
	(*person).name = "Mary"
}
