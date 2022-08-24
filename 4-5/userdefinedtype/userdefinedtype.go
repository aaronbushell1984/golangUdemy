package userdefinedtype

import (
	"fmt"
	"reflect"
)

func PrintType(mytype interface{}) string {
	return fmt.Sprintf("%T", mytype)
}

func PrintUnderlyingType(mytype interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(mytype).Kind())
}

type mytypeint int

var result int

// Conversion is idiom in go not casting
// User defined type is converted/reassigned here to its kind or unerlying type of int
func PrintUserDefinedIntConversion(myint mytypeint) string {
	result = int(myint)
	return fmt.Sprintf("%T", result)
}
