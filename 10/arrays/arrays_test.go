package arrays

import "fmt"

func ExampleCreateEmptyArrayOfSize() {
	fmt.Print(CreateEmptyArrayOfSize(5))
	fmt.Print(CreateEmptyArrayOfSize(10))
	// Output:
	// [0 0 0 0 0]
	// [0 0 0 0 0 0 0 0 0 0]
}

func ExampleCreateArrayOfSizeAndReassignOneValue() {
	fmt.Println(CreateArrayOfSizeAndReassignOneValue(5, 3, 100))
	fmt.Println(CreateArrayOfSizeAndReassignOneValue(2, 3, 100))
	fmt.Println(CreateArrayOfSizeAndReassignOneValue(-1, 3, 100))
	fmt.Println(CreateArrayOfSizeAndReassignOneValue(2, -3, 100))
	// Output:
	// [0 0 0 100 0]
	// There is no position 3 in an array which is 2 big
	// Size and position must be 1 or over
	// Size and position must be 1 or over
}

func ExampleGetLengthArray() {
	var array4 [4]int
	array3 := [3]int{1, 2, 3}
	fmt.Println(GetLengthArray(array4[:]))
	fmt.Println(GetLengthArray(array3[:]))
	// Output:
	// 4
	// 3
}
