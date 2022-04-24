package variables

import "fmt"

func GetReassignedVariables() int {
	// go language idiom
	// short declaration operator - "declares" and "assigns" variable
	x := 42
	fmt.Println(x)
	// assign new value
	x = 99
	return x
}

func GetExpressionVariables() int {
	// statements are usually oneline is one line and made of expressions (100 + 24 or 24 etc)
	// operator is + operands 100 and 24 in this example
	y := 100 + 24
	return y
}

func GetLocalVariables() string {
	// short declaration operator has local scope
	local := "Available in function"
	return local
}

var global = "Available outside function"

func GetGlobalVariables() string {
	return global
}
