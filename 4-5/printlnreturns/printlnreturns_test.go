package printlnreturns

import "fmt"

func ExamplePrintNumberAndError() {
	fmt.Println(PrintNumberAndError())
	// Output:
	// Capturing return values from fmt.Println:
	// 42 <nil>
}

func ExamplePrintNumberVoidError() {
	fmt.Println(PrintNumberVoidError())
	// Output:
	// Capturing return values from fmt.Println:
	// 42
}
