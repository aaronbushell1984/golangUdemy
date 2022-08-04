// Package recursion demonstrates a function which calls itself in go
//
// Recursion can often be acheived using a loop and can cause memory issues
package recursion

// GetRecursiveFactorial provides the factorial of a number using recursive:
//	return number * GetRecursiveFactorial(number -1)
// The loop must have a condition to stop infinite call:
//	if number == 0 {
// 		return 1
// 	}
func GetRecursiveFactorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * GetRecursiveFactorial(number -1)
}

// GetLoopFactorial provides the factorial of a number using a loop:
//	for ; number > 1; number-- {
// 		res *= number
// 	}
func GetLoopFactorial(number int) int {
	res := 1
	if number < 2 {
		return res
	}
	for ; number > 1; number-- {
		res *= number
	}
	return res
}