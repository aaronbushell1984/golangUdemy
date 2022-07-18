package spread

import "fmt"

func ExampleAppendSlice() {
	madeSlice := make([]int, 4, 4)
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(AppendSlice(madeSlice, 6, 9, 9, 1000))
	fmt.Println(AppendSlice(slice, slice...))
	// Output:
	// [0 0 0 0 6 9 9 1000]
	// [0 1 2 3 4 5 6 0 1 2 3 4 5 6]
}
