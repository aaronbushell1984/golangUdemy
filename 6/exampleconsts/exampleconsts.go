// Package exampleconsts demonstrates the use of const in go
package exampleconsts

import "fmt"

type PiConst struct {
	π float32
	pi float32
}

func ReturnTwoPiConsts() string {
	piConst := PiConst{
		π: 3.14,
		pi: π,
	}
	return fmt.Sprintf("%f and %f", π, piConst.pi)
}