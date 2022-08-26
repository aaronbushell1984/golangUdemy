// Package fuzzexample demonstrates use of a fuzz test in go
//
// Fuzzing the square root function reveals issues of imaginary numbers handled with:
//
//	if number < 0 {
//		number := math.Abs(number)
//		imaginary := math.Sqrt(number)
//		return err, imaginary, "i"
//	}
//
// Precision is revealed in fuzzing and handled by restricting SquareRoot:
//
//	if number > 999999999 {
//		return fmt.Errorf("number above maximum allowed"), 0, ""
//	}
//
// This error is skipped in fuzzing using:
//
//	if err != nil {
//		return
//	}
//
// The output is truncated for reasonable display and a higher precision used to ensure always accurate:
//
//	before := fmt.Sprintf("%.2f", math.Round(n*1000000)/1000000)
//	after := fmt.Sprintf("%.2f", math.Round(reSquared*1000000)/1000000)
package fuzzexample

import (
	"fmt"
	"math"
)

// SquareRoot uses:
//
//	math.Sqrt(number)
//
// to provide a square root
//
//	"number above maximum allowed"
//
// is returned as an error if a number over 999999999 is provided
//
//	"i"
//
// is returned as a string to signify imaginary number if provided number is negative
func SquareRoot(number float64) (error, float64, string) {
	var err error
	if number > 999999999 {
		return fmt.Errorf("number above maximum allowed"), 0, ""
	}
	if number < 0 {
		number := math.Abs(number)
		imaginary := math.Sqrt(number)
		return err, imaginary, "i"
	}
	return err, math.Sqrt(number), ""
}

// Square squares a number. If the number is imaginary then the result is negative
func Square(number float64, imaginary string) float64 {
	if imaginary == "i" {
		return -(number * number)
	}
	return number * number
}
