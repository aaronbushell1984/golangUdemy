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

// ReturnConstTrafficLights returns traffic light constants
func ReturnConstTrafficLights() string {
	const (
		light1 = "RED"
		light2 = "AMBER"
		light3 = "GREEN"
	)
	return fmt.Sprintf("%v, %v, %v", light1, light2, light3)
}
