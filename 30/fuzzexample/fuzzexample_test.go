package fuzzexample

import (
	"fmt"
	"math"
	"testing"
)

func ExampleSquareRoot() {
	_, root, _ := SquareRoot(16)
	fmt.Println(root)
	_, imaginaryBase, i := SquareRoot(-16)
	fmt.Println(imaginaryBase, i)
	// Output:
	// 4
	// 4 i
}

func ExampleSquare() {
	fmt.Println(Square(4, ""))
	fmt.Println(Square(4, "i"))
	// Output:
	// 16
	// -16
}

func FuzzSquareRoot(f *testing.F) {
	testcases := []float64{16, 4, 32}
	for _, v := range testcases {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, n float64) {
		var reSquared float64
		err, squareRoot, i := SquareRoot(n)
		if err != nil {
			return
		}
		if i == "i" {
			reSquared = Square(squareRoot, "i")
		} else {
			reSquared = Square(squareRoot, "")
		}
		before := fmt.Sprintf("%.2f", math.Round(n*1000000)/1000000)
		after := fmt.Sprintf("%.2f", math.Round(reSquared*1000000)/1000000)
		if before != after {
			t.Errorf("Before: %v After: %v, Error: %v", before, after, err)
		}
	})
}
