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

func ExampleGetFromSliceByIndex() {
	integers := []int{1, 2, 3}
	fmt.Println(GetFromSliceByIndex(integers, 2))
	fmt.Println(GetFromSliceByIndex(integers, 0))
	// Output:
	// 3
	// 1
}

func ExampleGetAllFromSlice() {
	integers := []int{1, 2, 3}
	moreIntegers := []int{-11, 2, 333, 9}
	fmt.Println(GetAllFromSlice(integers))
	fmt.Println(GetAllFromSlice(moreIntegers))
	// Output:
	// [1 2 3]
	// [-11 2 333 9]
}

func ExampleGetSliceOfSlice() {
	integersEasy := []int{0, 1, 2, 3, 4, 5}
	integersCaution := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(GetSliceOfSlice(integersEasy, 0, 4))
	fmt.Println(GetSliceOfSlice(integersCaution, 0, 4))
	// Output:
	// [0 1 2 3]
	// [1 2 3 4]
}

func ExampleGetAppendedSlices() {
	slice1 := []int{0, 1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(GetAppendedSlices(slice1, slice2))
	fmt.Println(GetAppendedSlices(slice2, slice1))
	// Output:
	// [0 1 2 3 4 5 1 2 3 4 5 6]
	// [1 2 3 4 5 6 0 1 2 3 4 5]
}

func ExampleGetDeletedSlice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(GetDeletedSlice(slice, 0, 2, 5, 6))
	// Output:
	// [0 1 2 5 6]
}
