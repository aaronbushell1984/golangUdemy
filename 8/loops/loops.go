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

// ReturnNumbersOneToTen returns the string representing slice of numbers from 1 to 10
//
// Demonstrates a single condition loop with:
//	for i <= 10 {
//		numbers = append(numbers, i)
//		i++
//	}
func ReturnNumbersOneToTen() string {
	i := 1
	var numbers []int
	for i <= 10 {
		numbers = append(numbers, i)
		i++
	}
	return fmt.Sprint(numbers)
}

// ReturnNumbersOneToTenWithBreak returns the string representing slice of numbers from 1 to 10
//
// Demonstrates an infinite loop with a break with:
//	for {
//		if i > 10 {
//			break
//		}
//		numbers = append(numbers, i)
//		i++
//	}
func ReturnNumbersOneToTenWithBreak() string {
	i := 1
	var numbers []int
	for {
		if i > 10 {
			break
		}
		numbers = append(numbers, i)
		i++
	}
	return fmt.Sprint(numbers)
}

// ReturnEvenNumbersWithContinue returns even numbers between two values
//
// "continue" is used to skip 0:
//	for i := start; i <= end; i++ {
//		if i == 0 {
//			continue
//		}
//		if i%2 == 0 {
//			evens = append(evens, i)
//		}
//	}
func ReturnEvenNumbersWithContinue(start int, end int) string {
	if start > end {
		return "End number must be greater or equal to start number"
	}
	var evens []int
	for i := start; i <= end; i++ {
		if i == 0 {
			continue
		}
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return fmt.Sprint(evens)
}

// ReturnAsciiNumbersAsCharacters returns ascii character representations of the numbers between supplied range
func ReturnAsciiNumbersAsCharacters(start int, end int) string {
	switch {
	case start < 33:
		return "There are no regular ascii characters below 33"
	case start > end:
		return "End number must be greater or equal to start number"
	}

	var ascii []int
	for i := start; i <= end; i++ {
		ascii = append(ascii, i)
	}
	return fmt.Sprintf("%c", ascii)
}
