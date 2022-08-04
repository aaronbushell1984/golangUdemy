package recursion

import (
	"fmt"
	"testing"
)

func ExampleGetRecursiveFactorial() {
	fmt.Println(GetRecursiveFactorial(1))
	fmt.Println(GetRecursiveFactorial(2))
	fmt.Println(GetRecursiveFactorial(3))
	fmt.Println(GetRecursiveFactorial(4))
	// Output:
	// 1
	// 2
	// 6
	// 24
}

func BenchmarkGetRecursiveFactorial(b *testing.B) {
	GetRecursiveFactorial(100)
}

func ExampleGetLoopFactorial() {
	fmt.Println(GetLoopFactorial(1))
	fmt.Println(GetLoopFactorial(2))
	fmt.Println(GetLoopFactorial(3))
	fmt.Println(GetLoopFactorial(4))
	// Output:
	// 1
	// 2
	// 6
	// 24
}

func BenchmarkGetLoopFactorial(b *testing.B) {
	GetLoopFactorial(100)
}



