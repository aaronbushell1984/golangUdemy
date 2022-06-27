// Package exampletypes demonstrates basic type in Go
package exampletypes

import (
	"fmt"
	"reflect"
)

// GetTypeReflect uses the reflect standard library to return the type of any given variable.
// 	reflect.TypeOf(variable)
// accesses the type of given variable.
//	 variable interface{}
// in method signature denotes any type as the variable argument.
func GetTypeReflect(variable interface{}) string {
	return reflect.TypeOf(variable).String()
}

// GetTypePercentT returns the type of any given variable.
// 	fmt.Sprintf(%T, variable)
//or
//	fmt.Sprintf(#{variable})
// returns the type of variable.
func GetTypePercentT(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}
