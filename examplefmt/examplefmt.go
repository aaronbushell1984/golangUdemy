package examplefmt

import (
	"fmt"
)

func PrintToConsole(variable interface{}) {
	fmt.Print("Printing without formatting:", variable)
}

func GetBinaryFromInt(number int) string {
	return fmt.Sprintf("%b", number)
}

func GetHexidecimalFromInt(number int) string {
	return fmt.Sprintf("%x", number)
}

func GetOctalHexidecimalFromInt(number int) string {
	return fmt.Sprintf("%#x", number)
}