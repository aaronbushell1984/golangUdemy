// Package table tests demonstrates tabular tests in go
package tabletests

// Sum returns the sum of any number of provided numbers
//
// Tabular test setup is - Create a test struct for data:
//	type test struct {
//		data 	[]int
//		answer 	int
//	}
// Declare a slice of test values containing the test data:
//	tests := []test{
//		{data: []int{1, 2, 3}, answer: 6},
//		{[]int{0, 2, 3}, 5},
//		{[]int{-2, 2, 3}, 3},
//	}
// Range over the data and check answers for errors:
//		for _, v := range tests {
//		got := Sum(v.data...)
//		if got != v.answer {
//			t.Errorf("Got: %v Want: %v", got, v.answer)
//		}
//	}
func Sum(numbers ...int) int {
	var res int
	for _, v := range numbers {
		res += v
	}
	return res
}
