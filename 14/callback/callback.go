// Package callback demonstrates passing a function into the argument of a function aka callback
package callback

// GetSum returns the total of any number of provided int values
func GetSum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

// GetEven returns only the even numbers from any number of provided numbers
func GetEven(numbers ...int) []int {
	var evens []int
	for _, v := range numbers {
		if v%2 == 0 && v != 0 {
			evens = append(evens, v)
		}
	}
	return evens
}

// GetEvenSum returns to sum of even numbers from any given numbers using a callback:
//	GetEvenSum(getSum func(numbers ...int) int, numbers ...int) int
// Any function which takes any number of int and returns an int can be passed:
//	numbers := []int{1, 2, 3, 4, 5}
//	fmt.Println(GetEvenSum(GetSum, numbers...))
// In this argument GetSum is passed without ()
//
// The getSum is taken from argument and not GetSum() from package:
//	return getSum(GetEven(numbers...)...)
// GetEven is taken from package and []int spread to pass to callback getSum()
func GetEvenSum(getSum func(numbers ...int) int, numbers ...int) int {
	return getSum(GetEven(numbers...)...)
}

// GetOdd returns only the odd numbers from any number of provided numbers
func GetOdd(numbers ...int) []int {
	var odds []int
	for _, v := range numbers {
		if v%2 != 0 && v != 0 {
			odds = append(odds, v)
		}
	}
	return odds
}

// GetOddSum returns the sum of the odd numbers provided
//
// It uses GetSum() and GetOdd() in  the pacakge package to demonstrate not using a callback
func GetOddSum(numbers ...int) int {
	return GetSum(GetOdd(numbers...)...)
}
