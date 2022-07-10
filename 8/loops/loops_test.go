package loops

import (
	"fmt"
	"testing"
)

func ExamplePrintNumbersLoop() {
	fmt.Println(PrintNumbersLoop(0, 1))
	fmt.Println(PrintNumbersLoop(0, 10))
	fmt.Println(PrintNumbersLoop(-10, 10))
	// Output:
	// [0 1]
	// [0 1 2 3 4 5 6 7 8 9 10]
	// [-10 -9 -8 -7 -6 -5 -4 -3 -2 -1 0 1 2 3 4 5 6 7 8 9 10]
}

func BenchmarkPrintNumbersLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintNumbersLoop(-10, 10)
	}
}

func ExampleReturnAllTwoBitOptions() {
	fmt.Println(ReturnAllTwoBitOptions())
	// Output:
	// [00 01 10 11]
}

func BenchmarkReturnAllTwoBitOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnAllTwoBitOptions()
	}
}
