// Package slices demonstrates the use of slices in go
package slices

import (
	"fmt"
	"reflect"
)

// ReturnSlice takes in any type and returns it if a slice or provides warning
func ReturnSlice(compLiteral interface{}) string {
	typ := reflect.ValueOf(compLiteral)
	switch {
	case typ.Kind() != reflect.Slice:
		return "This is not a slice"
	}
	return fmt.Sprintf("%v", compLiteral)
}

// GetFromSliceByIndex returns the integer from the provided slice at the index provided
func GetFromSliceByIndex(sliceOfInt []int, index int) int {
	return sliceOfInt[index]
}

// GetAllFromSlice returns all integers from slice as a string
//
// Demonstrates use of a "for range" which discards index:
//	for _, v := range sliceOfInt {
//		result = append(result, v)
//	}
func GetAllFromSlice(sliceOfInt []int) string {
	var result []int
	for _, v := range sliceOfInt {
		result = append(result, v)
	}
	return fmt.Sprint(result)
}

// GetSliceOfSlice returns a portion of a slice
//
// N.B. Indexes start at 0 -- AND
//	[start:startOfEnd]
// may not be intuitive
func GetSliceOfSlice(slice []int, start int, startOfEnd int) []int {
	return slice[start:startOfEnd]
}

// GetAppendedSlices returns the second slice appended to the end of the first
//
// Uses:
//	...
// Which unpacks slice2:
//	append(slice1, slice2...)
// Not be confused with:
//	...T
// Which in a function signature is a variadic parameter meaning, it can accept any number of arguments of the given type.
func GetAppendedSlices(slice1 []int, slice2 []int) []int {
	return append(slice1, slice2...)
}

func GetDeletedSlice(slice []int, start1 int, end1 int, start2 int, end2 int) []int {
	return append(slice[start1:end1+1], slice[start2:end2+1]...)
}
