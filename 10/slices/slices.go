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
//
//	for _, v := range sliceOfInt {
//		result = append(result, v)
//	}
func GetAllFromSlice(sliceOfInt []int) string {
	var result []int
	result = append(result, sliceOfInt...)
	return fmt.Sprint(result)
}

// GetSliceOfSlice returns a portion of a slice
//
// N.B. Indexes start at 0 -- AND
//
//	[start:startOfEnd]
//
// may not be intuitive
func GetSliceOfSlice(slice []int, start int, startOfEnd int) []int {
	return slice[start:startOfEnd]
}

// GetAppendedSlices returns the second slice appended to the end of the first
//
// Uses:
//
//	...
//
// Which unpacks slice2:
//
//	append(slice1, slice2...)
//
// Not be confused with:
//
//	...T
//
// Which in a function signature is a variadic parameter meaning, it can accept any number of arguments of the given type.
func GetAppendedSlices(slice1 []int, slice2 []int) []int {
	return append(slice1, slice2...)
}

// GetDeletedSlice returns two parts of a slice appended together
//
// Put another way, this deletes a part of a slice
func GetDeletedSlice(slice []int, start1 int, end1 int, start2 int, end2 int) []int {
	return append(slice[start1:end1+1], slice[start2:end2+1]...)
}

// GetCapacityOfSlice returns the capacity of the underlying array a slice is built on
func GetCapacityOfSlice(slice []int) int {
	return cap(slice)
}

// GetLengthOfSlice returns the number of items currently in a slice
func GetLengthOfSlice(slice []int) int {
	return len(slice)
}

// GetAppendedSlice returns a slice appended with any number of added int values
//
//	func BenchmarkGetAppendedSlice(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			slice := []int{0, 1, 2, 3, 4, 5, 6}
//			GetAppendedSlice(slice, 6, 9, 9, 1000)
//		}
//	}
//
// Runs in 54.37 ns/op
//
//	func BenchmarkGetAppendedMadeSlice(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			madeSlice := make([]int, 7, 11)
//			GetAppendedSlice(madeSlice, 6, 9, 9, 1000)
//		}
//	}
//
// Runs in 0.3670 ns/op
//
// # Appending to a slice is possible
//
// Making a slice with enough capacity for later appends is faster
func GetAppendedSlice(slice []int, number ...int) []int {
	return append(slice, number...)
}

// GetCombinedMultiDimensionalSlices combines any number of slices into a multi-dimensional slice
//
// Generic V is used to allow []string or []int:
//
//	GetCombinedMultiDimensionalSlices[V []string | []int]
//
// Variadic parameter ... allows any number of the type V:
//
//	(slice ...V)
//
// Return type is set to a slice of V which is a slice and therefore [][] (multi-dimensional)
//
//	[]V
func GetCombinedMultiDimensionalSlices[V []string | []int](slice ...V) []V {
	return slice
}

// DataIsInMap returns true if the key returns a result in provided map
func DataIsInMap(mapToTest map[int]string, key int) bool {
	if _, ok := mapToTest[key]; ok {
		return true
	}
	return false
}

// GetAsciiGlyphByDecimal returns the ascii glyph of provided decimal
func GetAsciiGlyphByDecimal(decimal int) string {
	return fmt.Sprintf("%c", decimal)
}

// GetMapOfAsciiGlyphsByDecimal returns ascii glyphs by decimal in a map
func GetMapOfAsciiGlyphsByDecimal(decimal ...int) map[int]string {
	m := make(map[int]string)
	for _, v := range decimal {
		m[v] = fmt.Sprintf("%c", v)
	}
	return m
}

// DeleteFromMapByKey deletes an entry by key from a map if it is present or does nothing
func DeleteFromMapByKey(intStringMap map[int]string, key int) {
	if DataIsInMap(intStringMap, key) {
		delete(intStringMap, key)
	}
}
