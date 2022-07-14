package slices

import "fmt"

func ExampleReturnSlice() {
	strings := []string{"Slices", "are", "composite", "literals", "and hold values of same type"}
	integers := []int{1, 2, 3}
	notSlice := 1
	fmt.Println(ReturnSlice(strings))
	fmt.Println(ReturnSlice(integers))
	fmt.Println(ReturnSlice(notSlice))
	// Output:
	// [Slices are composite literals and hold values of same type]
	// [1 2 3]
	// This is not a slice
}
