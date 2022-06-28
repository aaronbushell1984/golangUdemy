// Package stringtypes demonstrates string types in Go.
package stringtypes

import (
	"fmt"
	"strings"
)

// PrintStringAndType returns the type of any give string in a sentence.
func PrintStringAndType(name string) string {
	return name + " is a " + fmt.Sprintf("%T", name)
}

// ConvertStringToSliceOfByteString returns a slice of the bytes making up a given string
func ConvertStringToSliceOfByteString(s string) string {
	byteslice := []byte(s)
	return fmt.Sprintf("%v", byteslice)
}

// NewlineAndTabRawStringLiteral demonstrates using the ` characters to interpret a carriage return and tabs.
// 	`This is a:
//				raw string literal with new line and tabs`
// is returned as a raw string literal
func NewlineAndTabRawStringLiteral() string {
	return `This is a:
				raw string literal with new line and tabs`
}

// PrintStringCodepoints returns the codepoints of a given string as a string.
func PrintStringCodepoints(s string) string {
	var codepoints []string
	for i := 0; i < len(s); i++ {
		cp := fmt.Sprintf("%#U ", s[i])
		codepoints = append(codepoints, cp)
	}
	cps := strings.Join(codepoints, "")
	return cps
}
