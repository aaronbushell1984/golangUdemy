// Package loops demonstrates the use of loops in go
package loops

// PrintNumbersLoop returns all numbers between start and end inputs using syntax:
//	for initialize; condition; post
func PrintNumbersLoop(start int, end int) []int {
	var numbers []int
	for i := start; i <= end; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}
