// Package reflectexample demonstrates the use of reflect deep equal to compare slices:
//
//	!reflect.DeepEqual(got, want)
//
// NB. This does allow comparison of non type safe variables:
//
//	!reflect.DeepEqual(imAString, imAnIntSlice)
//
// The test package also shows a reminder of multiple tests using T.run:
//
//	t.Run("sum of tail of slices", func(t *testing.T) {
//
// Assigning a function to a variable in test function keeps the function scope to withing the test:
//
//	checkSums := func(t *testing.T, got []int, want []int) {
//		if !reflect.DeepEqual(got, want) {
//			t.Errorf("SumAllTails() Got: %v Want %v", got, want)
//		}
//	}
package reflectexample

// Sum returns the sum of values in the provided slice of integers
func Sum(numbers []int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}

// SumAll adds all the values in any number of provided slice of integers and returns a new slice of the totals
func SumAll(ns ...[]int) []int {
	var sums []int
	for _, v := range ns {
		sums = append(sums, Sum(v))
	}
	return sums
}

// SumAllTails adds all but the first values in any number of provided slice of integers and returns a new slice of the totals
func SumAllTails(ns ...[]int) []int {
	var sums []int
	for _, slice := range ns {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
