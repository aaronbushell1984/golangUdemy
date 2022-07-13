// Package exercisetwo is the third set of exercises from the course
package exercisethree

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

// ReturnYearsAlive returns all the years between supplied year and 1984
func ReturnYearsAlive(year int) string {
	result := "1984"
	if year < 1984 {
		return "Enter a year on or after " + result
	}
	for i := 1985; i <= year; i++ {
		result += fmt.Sprintf(" %v", i)
	}
	return result
}

// ReturnDivideFourRemainder finds the remainder of all numbers between the range when quartered
func ReturnDivideFourRemainder(start int, end int) []int {
	var remainders []int
	switch {
	case start > end:
		return remainders
	}
	for i := start; i <= end; i++ {
		k := i%4
		remainders = append(remainders, k)
	}
	return remainders
}

// GetResponseToFaveSport returns a response based on sport added using a switch statement
func GetResponseToFaveSport(faveSport string) string {
	switch faveSport {
	case "Football":
		return "World's most popular sport!"
	case "Tennis":
		return "Strawberries and cream!"
	case "Golf":
		return "Easy tiger!"
	}
	return "Not heard of that one!"
}