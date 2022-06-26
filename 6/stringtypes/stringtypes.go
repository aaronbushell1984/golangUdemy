package stringtypes

import (
	"fmt"
	"strings"
)

func PrintStringAndType(name string) string {
	return name + " is a " + fmt.Sprintf("%T", name)
}

func ConvertStringToSliceOfByteString(s string) string {
	byteslice := []byte(s)
	return fmt.Sprintf("%v", byteslice)
}

func NewlineAndTabRawStringLiteral() string {
	return `This is a:
				raw string literal with new line and tabs`
}

func PrintStringCodepoints(s string) string {
	var codepoints []string
	for i := 0; i < len(s); i++ {
		cp := fmt.Sprintf("%#U ", s[i])
		codepoints = append(codepoints, cp)
	}
	cps := strings.Join(codepoints, "")
	return cps
}
