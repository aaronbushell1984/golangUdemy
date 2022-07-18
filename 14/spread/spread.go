// Package spread demonstrates the use of spreading slices into individual values. Called spread operator in Javasript, unpack in IntelliJ and unfurling in course.
package spread

// AppendSlice takes any number of int (or a slice) and appends it to the given slice
//
// Spread operator, unpacking or unfurling happens when the ints are appended:
//	return append(slice, number...)
func AppendSlice(slice []int, number ...int) []int {
	return append(slice, number...)
}
