package deferexample

import (
	"fmt"
)

func ExampleAddFirst() {
	AddFirst()
	fmt.Println(result)
	// Output:
	// [first]
}

func ExampleAddSecond() {
	AddSecond()
	fmt.Println(result)
	// Output:
	// [second]
}

func ExampleAppendStrings() {
	fmt.Println(AppendStrings())
	// Output:
	// [first then second]
}

func ExampleDeferAppendStrings() {
	fmt.Println(DeferAppendStrings())
	// Output:
	// [then second]
}
