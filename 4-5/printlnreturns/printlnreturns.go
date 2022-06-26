// Package printlnreturns demonstrates return types using fmt.Println as an example
package printlnreturns

import "fmt"

// PrintNumberAndError prints n and err from fmt.Println
//
// fmt.Println has method signature:
//		func Println(a ...any) (n int, err error)
//
// therefore returns (n int, err error)
// 		n, err := fmt.Println("Capturing return values from fmt.Println:")
//
// captures the Number and Error
func PrintNumberAndError() string {
	n, err := fmt.Println("Capturing return values from fmt.Println:")
	return fmt.Sprintln(n, err)
}

// PrintNumberVoidError discards err by voiding with _
//
// fmt.Println has method signature:
//		func Println(a ...any) (n int, err error)
//
// the error could be discarded with _ (blank identifier)
//		n, _ := fmt.Println("Capturing return values from fmt.Println:")
func PrintNumberVoidError() string {
	n, _ := fmt.Println("Capturing return values from fmt.Println:")
	return fmt.Sprintln(n)
	//fmt.Println(n, err) this no throws compile error because err is _
}
