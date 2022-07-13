// Package arrays demonstrates the use of arrays in go
package arrays

import "fmt"

// CreateEmptyArrayOfSize creates an array of given size with default values of 0
func CreateEmptyArrayOfSize(size int) string {
	emptyArray := make([]int,size)
	return fmt.Sprintf("%v\n", emptyArray)
}

// CreateArrayOfSizeAndReassignOneValue returns an array string with declared size
//
// At the postion specified the value is added to array
func CreateArrayOfSizeAndReassignOneValue(size int, position int, value int) string {
	switch {
	case size < 1 || position < 1:
		return "Size and position must be 1 or over"
	case position > size-1:
		return fmt.Sprintf("There is no position %v in an array which is %v big", position, size)
	}
	emptyArray := make([]int,size)
	emptyArray[position] = value
	return fmt.Sprintf("%v", emptyArray)
}

