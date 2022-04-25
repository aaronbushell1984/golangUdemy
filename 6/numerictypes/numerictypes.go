package numerictypes

import "fmt"

func PrintType(num interface{}) string {
	return fmt.Sprintf("%T", num)
}