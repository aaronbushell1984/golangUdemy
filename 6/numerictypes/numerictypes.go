package numerictypes

import "fmt"

func PrintDefaultType(num interface{}) string {
	return fmt.Sprintf("%T", num)
}