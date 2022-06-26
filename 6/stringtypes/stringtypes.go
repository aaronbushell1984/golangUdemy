package stringtypes

import "fmt"

func PrintStringAndType(name string) string {
	return name + " is a " + fmt.Sprintf("%T", name)
}

func ConvertStringToSliceOfByteString(s string) string {
	byteslice := []byte(s)
	return fmt.Sprintf("%v", byteslice)
}

func NewlineAndTabRawStringLiteral() string {
	return `"Mrunalini is a:
				string"`
}

//func PrintStringCodepoints(s string) []int {
//	var codepoints []int
//	for i := 0; i > len(s); i++ {
//		codepoints = append(codepoints, s)
//	}
//	return codepoints
//}
