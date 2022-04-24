package examplefmt

import (
	"fmt"
	"io"
)

// Capture print in io for testing
// Use Fprint() with no formatting
func PrintConsoleAndSaveToBuffer(w io.Writer, variable interface{}) {
	fmt.Fprint(w, "Printing without formatting: ", variable)
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