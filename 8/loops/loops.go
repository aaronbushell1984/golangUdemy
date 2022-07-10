// Package loops demonstrates the use of loops in go
package loops

import (
	"fmt"
)

// PrintNumbersLoop returns all numbers between start and end inputs using syntax:
//	for initialize; condition; post
func PrintNumbersLoop(start int, end int) []int {
	var numbers []int
	for i := start; i <= end; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

// ReturnAllTwoBitOptions returns all binary options when using two bits
func ReturnAllTwoBitOptions() []string {
	var numbers []string
	for i := 0; i <= 1; i++ {
		for k := 0; k <= 1; k++ {
			binary := fmt.Sprintf("%d%d", i, k)
			numbers = append(numbers, binary)
		}
	}
	return numbers
}
