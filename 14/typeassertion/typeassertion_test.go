package typeassertion

import "fmt"

func ExampleIsString() {
	s := "I'm a string"
	n := 1
	fmt.Println(IsString(s))
	fmt.Println(IsString(n))
	// Output:
	// true
	// false
}

func ExampleIsIntOrInt64() {
	s := "I'm a string"
	n := 1
	var sf int64 = 64
	fmt.Println(IsIntOrInt64(s))
	fmt.Println(IsIntOrInt64(n))
	fmt.Println(IsIntOrInt64(sf))
	// Output:
	// false
	// true
	// true
}
