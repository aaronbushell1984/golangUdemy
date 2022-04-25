package examplefmt

import (
	"fmt"
	"io"
)

// Capture print in io for testing
// Use Fprint() without formatting
func PrintConsoleAndSaveToBuffer(w io.Writer, variable interface{}) {
	description := "Printing without formatting: "
	fmt.Fprint(w, description, variable)
}

func PrintfConsoleAndSaveToBuffer(w io.Writer, variable interface{}) {
	description := "Printing with formatting: "
	fmt.Fprintf(w, "%s\n\t%s", description, variable)
}

func GetBinaryFromInt(number int) string {
	return fmt.Sprintf("%b", number)
}

func GetHexidecimalFromInt(number int) string {
	return fmt.Sprintf("%x", number)
}

func GetOctalHexidecimalFromInt(number int) string {
	return fmt.Sprintf("%#x", number)
}