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

func ExampleFirstInLastOut() {
	fmt.Println(FirstInLastOut())
	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
}

func ExampleDeferCapturedAtTimeOfCall() {
	fmt.Println(DeferCapturedAtTimeOfCall())
	// Output:
	// 0
	// 1
}
