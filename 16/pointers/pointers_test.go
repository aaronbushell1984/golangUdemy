package pointers

import "fmt"

func ExampleGetAddressOfValue() {
	value := "Value"
	fmt.Printf("%T", GetAddressOfValue(value))
	// Output:
	// *interface {}
}

func ExampleAssignToAddress() {
	ban := "banana"
	fmt.Printf("%T\n", ban)
	fmt.Printf("%T", AssignToAddress(ban))
	// Output:
	// string
	// string
}