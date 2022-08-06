package deferexample

import (
	"fmt"
)

func ExampleGetFirstSecond() {
	fmt.Println(GetFirstSecond())
	// Output:
	// [first second]
}

func ExampleGetSecondFirst() {
	fmt.Println(GetSecondFirst())
	// Output:
	// [second]
}