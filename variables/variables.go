package main 

import "fmt"

func main() {
	// go language idiom
	// short declaration operator - "declares" and "assigns" variable
	x := 42
	fmt.Println(x)
	// assign new value
	x = 99
	fmt.Println(x)
	// statements are usually oneline is one line and made of expressions (100 + 24 or 24 etc)
	// operator is + operands 100 and 24 in this example
	y := 100 + 24
	fmt.Println(y)
}