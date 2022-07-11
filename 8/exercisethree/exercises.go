// Package exercisetwo is the third set of exercises from the course
package exercises

import "fmt"

// ReturnOneToTenThousand returns numbers 1-10,000
func ReturnOneToTenThousand() string {
	var numbers []int
	for i := 1; i <= 10000; i++ {
		numbers = append(numbers, i)
	}
	return fmt.Sprintf("%v", numbers)
}

// ReturnUpperAlphabetRuneCodePointsThreeTimes returns capital characters A-Z, thrice, in format:
//	U+0041 'A'
// 	U+0041 'A'
//	U+0041 'A'
func ReturnUpperAlphabetRuneCodePointsThreeTimes() string {
	result := ""
	for i := 65; i <= 90; i++ {
		codepoint := fmt.Sprintf(`%#U
`, i)
		for i := 0; i < 3; i++ {
			result += codepoint
		}
	}
	return fmt.Sprintf("%v", result)
}
