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