// Package exercisefour is the fourth set of exercises from the course
package exercisefour

import "fmt"

// GetTypeOfArray returns the type of provided array as a string
func GetTypeOfArray(arr interface{}) string {
	return fmt.Sprintf("%T", arr)
}

// GetTypeOfSlice returns the type of provided slice as a string
//
// Four slice types can be accepted which is declared using:
//	[T []int | []string | []float64 | []float32]
func GetTypeOfSlice[T []int | []string | []float64 | []float32](slice T) string {
	return fmt.Sprintf("%T", slice)
}

// GetPartSlice returns selected part of a slice
func GetPartSlice(slice []int, start int, end int) []int {
	return append(slice[start : end+1])
}

// AddToSlice returns a slice of integers appended by any number of integers
func AddToSlice(intSlice []int, integer ...int) []int {
	return append(intSlice, integer...)
}

// DeleteFromSliceByIndex deletes a section of a slice starting and finishing at provided indices
func DeleteFromSliceByIndex(slice []int, start int, end int) []int {
	return append(slice[:start], slice[end:]...)
}

// GetSliceCapacity returns the capacity of underlying array of slice
func GetSliceCapacity(sliceString []string) int {
	return cap(sliceString)
}

// GetSliceSize returns the size (or length) of slice
func GetSliceSize(sliceString []string) int {
	return len(sliceString)
}

// GetIndexedMapFromSlice takes a slice of strings and returns a map with the values indexed
//
// Capacity is used to define the size of map for efficiency
func GetIndexedMapFromSlice(sliceString []string, capacity int) map[int]string {
	m := make(map[int]string, capacity)
	for i := 0; i < len(sliceString); i++ {
		m[i] = sliceString[i]
	}
	return m
}

// GetMultiDimensionalSlicesFromSlices combines any number of slices into a multi-dimensional slice
//
// Generic V is used to allow []string or []int:
//	GetCombinedMultiDimensionalSlices[V []string | []int]
// Variadic parameter ... allows any number of the type V:
//	(slice ...V)
// Return type is set to a slice of V which is a slice and therefore [][] (multi-dimensional)
//	[]V
func GetMultiDimensionalSlicesFromSlices[V []string | []int](slice ...V) []V {
	return slice
}

// AddRecordToMap adds the provided key and data to provided map
func AddRecordToMap(mapStringString map[string]string, key string, data string) map[string]string {
	mapStringString[key] = data
	return mapStringString
}

// DataIsInMap returns true if the key returns a result in provided map
func DataIsInMap(mapToTest map[string]string, key string) bool {
	if _, ok := mapToTest[key]; ok {
		return true
	}
	return false
}

// DeleteFromMapByKey deletes an entry by key from a map if it is present or does nothing
func DeleteFromMapByKey(mapStringString map[string]string, key string) {
	if DataIsInMap(mapStringString, key) {
		delete(mapStringString, key)
	}
}
