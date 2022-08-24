package exercises

import "fmt"

func GetShortDeclarations() string {
	x := 42
	y := "James Bond"
	z := true
	return fmt.Sprintf("%v %v %v", x, y, z)
}

var x int
var y string
var z bool

func GetVarDeclarationZeroValues() string {
	return fmt.Sprintf("%v%v %v", x, y, z)
}

var num = 42
var jb = "James Bond"
var check = true

func GetVarDeclarationsAsString() string {
	s := fmt.Sprintf("%v %v %v", num, jb, check)
	return s
}

func PrintValueOfUserDefinedType(myvar interface{}) string {
	return fmt.Sprint(myvar)
}

func PrintTypeOfUserDefinedType(myvar interface{}) string {
	return fmt.Sprintf("%T", myvar)
}
