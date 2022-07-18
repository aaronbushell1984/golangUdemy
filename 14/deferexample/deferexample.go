// Package deferexample demonstartes the use of defer in go
package deferexample

var result []string

// AddFirst adds "first" to the package scope variable:
//	var result []string
// This is setup for AppendStrings() Vs DeferAppendStrings()
func AddFirst() {
	result = append(result, "first")
}

// AddSecond adds "second" to the package scope variable:
//	var result []string
// This is setup for AppendStrings() Vs DeferAppendStrings()
func AddSecond() {
	result = append(result, "second")
}

// AppendStrings demonstrates using AddFirst() and AddSecond() in order for comparison with DeferAppendString()
//
// The output with "then" added is:
//	[first then second]
func AppendStrings() []string {
	AddFirst()
	result = append(result, "then")
	AddSecond()
	return result
}

// DeferAppendStrings demonstrates deferring a function until end of execution block
//
// AddFirst() executes after the return statement:
//	defer AddFirst()
//	result = append(result, "then")
//	AddSecond()
//	return result
// Therefore the output is:
//	[then second]
func DeferAppendStrings() []string {
	defer AddFirst()
	result = append(result, "then")
	AddSecond()
	return result
}
