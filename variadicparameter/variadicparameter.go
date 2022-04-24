package variadicparameter

import "fmt"

// from fmt docs
// func Println(a ...any) (n int, err error) - takes any number of any types of arguments (variadic parameter)
func VaridaicParameterPrint() {
	fmt.Println("fmt.Prinln can accept any number of different type variable such as this string and following number and boolean:", 42, true)
}
