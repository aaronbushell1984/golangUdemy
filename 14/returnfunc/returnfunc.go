// Package returnfunc demonstrates returning a function from a function
package returnfunc

// GetIntViaFunc returns a function which returns an int.
//	func GetIntViaFunc() func() int {
//		return func() int {
//			return 42
//		}
//	}
// Capturing the function expression:
//	theFunc := GetIntViaFunc()
// Allows the result to be captured from the function call:
//	answer := theFunc()
func GetIntViaFunc() func() int {
	return func() int {
		return 42
	}
}