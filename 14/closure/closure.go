// Package closure demonstrates closures in go
package closure

// Increment returns a function which increments an int by 1
//
// This demonstrates closure of the x variable which is closed by the anonymous function:
//	func Increment() func() int {
//		x := 0
//		return func() int {
// 			x++
// 			return x
// 		}
// 	}
// As this is closed each different function expression can increment independently; see example
func Increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}