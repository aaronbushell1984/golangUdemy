// Package anonymousfunctions demonstrates anonymous functions in go
package anonymousfunctions

// AnonymousFunction contains an un-named self executing function:
//	func() {
//		result = "I was called by an anonymous function inside a named one"
//	}()
// Paradoxically it is contained in a named function to allow the example test
func AnonymousFunction() string {
	result := ""
	func() {
		result = "I was called by an anonymous function inside a named one"
	}()
	return result
}
