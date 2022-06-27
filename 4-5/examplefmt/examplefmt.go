// Package examplefmt demonstrates various uses of the fmt standard library
package examplefmt

import (
	"fmt"
	"io"
)

// PrintConsoleAndSaveToBuffer uses io.Writer and accepts, prints and saves any type of variable.
// 	fmt.Fprint
// prints to the console without allowing formatting
// 	io.Writer
// captures the output which is useful for testing.
func PrintConsoleAndSaveToBuffer(w io.Writer, variable interface{}) {
	description := "Printing without formatting: "
	fmt.Fprint(w, description, variable)
}

// PrintfConsoleAndSaveToBuffer uses io.Writer and accepts, prints and saves any type of variable.
// 	fmt.Fprintf
// prints to the console and allows formatting using verbs like %T
// 	io.Writer
// captures the output which is useful for testing.
func PrintfConsoleAndSaveToBuffer(w io.Writer, variable interface{}) {
	description := "Printing with formatting:"
	fmt.Fprintf(w, "%s\n\t%s", description, variable)
}

// GetBinaryFromInt returns the binary from a given int.
// 	fmt.Sprintf
// allows print to be returned as a string
// 	%b
// is the verb for binary
func GetBinaryFromInt(number int) string {
	return fmt.Sprintf("%b", number)
}

// GetHexidecimalFromInt returns the hexidecimal from a given int.
// 	fmt.Sprintf
// allows print to be returned as a string
// 	%x
// is the verb for hexidecimal
func GetHexidecimalFromInt(number int) string {
	return fmt.Sprintf("%x", number)
}

// GetOctalHexidecimalFromInt returns the octo hexidecimal from a given int.
// 	fmt.Sprintf
// allows print to be returned as a string
// 	%#x
// is the verb for hexidecimal
func GetOctalHexidecimalFromInt(number int) string {
	return fmt.Sprintf("%#x", number)
}
