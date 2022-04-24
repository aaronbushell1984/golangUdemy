package exampletypes

import (
	"fmt"
	"reflect"
)

// reflect package contains TypeOf to return type
// variable interface{} allows any type into function
func GetTypeReflect(variable interface{}) string {
	return reflect.TypeOf(variable).String()
}

// %T and Printf on fmt package can print type to console
// %T and Sprintf on fmt package can return type
func GetTypePercentT(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}