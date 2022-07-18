// Package variadicparameter demonstrates the use of a variadic parameter in go
package variadicparameter

// CombineSlices takes any number of slices of int or string and returns a multi-dimensional slice
//
// Variadic parameter "..." allows this:
//	CombineSlices[V []string | []int](slices ...V) []V
func CombineSlices[V []string | []int](slices ...V) []V {
	return slices
}
