// Package funcexpression demonstrates assigning a function to a value
package funcexpression

// GetHelloWorld assigns a function to value HelloWorld:
//
//	HelloWorld := func() string {
//		return fmt.Sprintf("Hello World")
//	}
//
// This then returns this function which in turn was passed "Hello World":
//
//	return Hello()
func GetHelloWorld() string {
	HelloWorld := func() string {
		return "Hello World"
	}
	return HelloWorld()
}
