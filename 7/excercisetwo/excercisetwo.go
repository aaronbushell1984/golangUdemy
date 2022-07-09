// Package excercisetwo is the second set of excercises from the course
package excercisetwo

import "fmt"

// ReturnDecimalBinaryHex takes a number and returns the decimal, binary and hexidecimal format
func ReturnDecimalBinaryHex(i int) string {
	return fmt.Sprintf("Decimal: %v Binary: %b Hex: %#x", i, i, i)
}

// CheckEquals checks two integers to see if they are equal
func CheckEquals(i int, i2 int) bool {
	switch {
	case i == i2:
		return true
	default:
		return false
	}
}

// CheckLessOrEqual checks two integers to see if the first is less than or equal to second
func CheckLessOrEqual(i int, i2 int) bool {
	switch {
	case i <= i2:
		return true
	default:
		return false
	}
}

// CheckGreaterOrEqual checks two integers to see if the first is greater than or equal to second
func CheckGreaterOrEqual(i int, i2 int) bool {
	switch {
	case i >= i2:
		return true
	default:
		return false
	}
}

// CheckNotEqual checks two integers to see if the first is not equal to second
func CheckNotEqual(i int, i2 int) bool {
	switch {
	case i != i2:
		return true
	default:
		return false
	}
}

// CheckLessThan checks two integers to see if the first is less than second
func CheckLessThan(i int, i2 int) bool {
	switch {
	case i < i2:
		return true
	default:
		return false
	}
}

// CheckGreaterThan checks two integers to see if the first is greater than second
func CheckGreaterThan(i int, i2 int) bool {
	switch {
	case i > i2:
		return true
	default:
		return false
	}
}

// ReturnConstTrafficLights returns traffic light constants.
//
// The first light is a constant of a kind and the second two are not:
//	const (
//		light1 string = "RED"
//		light2 = "AMBER"
//		light3 = "GREEN"
//	)
func ReturnConstTrafficLights() string {
	const (
		light1 string = "RED"
		light2        = "AMBER"
		light3        = "GREEN"
	)
	return fmt.Sprintf("%v, %v, %v", light1, light2, light3)
}

// PrintIntBeforeAndAfterBitShift takes an int and returns the decimal, binary and hex before and after bit shift
func PrintIntBeforeAndAfterBitShift(i int) string {
	before := fmt.Sprintf("Before: %v %b %#x", i, i, i)
	i = i << 1
	after := fmt.Sprintf("After: %v %b %#x", i, i, i)
	return before + ", " + after
}

// MaintainStringFormatting returns the same passed in string.
//
// The test demonstrates it maintains raw string literal:
//	func TestMaintainStringFormatting(t *testing.T) {
//		want := "This maintains formatting using a\nRaw\nString\nLiteral"
//		example := `This maintains formatting using a
//	Raw
//	String
//	Literal`
//		if got := MaintainStringFormatting(example); got != want {
//			t.Errorf("MaintainStringFormatting(): got: %q want: %q", got, want)
//		}
//	}
func MaintainStringFormatting(s string) string {
	return fmt.Sprintf("%v", s)
}

// ReturnLastFourYears takes the year and returns the previous four using iota
func ReturnLastFourYears(year int) string {
	const (
		_ = iota
		y1
		y2
		y3
		y4
	)
	year1 := year - y1
	year2 := year - y2
	year3 := year - y3
	year4 := year - y4
	return fmt.Sprintf("%v, %v, %v, %v", year1, year2, year3, year4)
}
