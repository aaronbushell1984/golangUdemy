package exerciseseven

import "fmt"

func ExampleReturnAddressType() {
	banana := "banana"
	fmt.Println(ReturnAddressType(banana))
	// Output:
	// *string
}

func ExampleChangeMe() {
	jb := person{
		name: "James",
	}
	fmt.Println(jb)
	ChangeMe(&jb)
	fmt.Println(jb)
	// Output:
	// {James}
	// {Mary}
}
