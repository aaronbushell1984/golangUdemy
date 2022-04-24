package examplefmt

import (
	"fmt"
)

func GetBinaryFromInt(number int) string {
	return fmt.Sprintf("%b", number)
}

