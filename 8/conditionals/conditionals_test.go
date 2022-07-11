package conditionals

import (
	"fmt"
	"testing"
)

func ExampleReturnSuppliedNameIfOverThreeCharactersCase() {
	fmt.Println(ReturnSuppliedNameIfOverThreeCharactersCase("Mrunalini"))
	fmt.Println(ReturnSuppliedNameIfOverThreeCharactersCase("Mru"))
	fmt.Println(ReturnSuppliedNameIfOverThreeCharactersCase("Aaron"))
	// Output:
	// Mrunalini
	// Name must be over three characters
	// Aaron
}

func BenchmarkReturnSuppliedNameIfOverThreeCharactersCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnSuppliedNameIfOverThreeCharactersCase("Mrunalini")
		ReturnSuppliedNameIfOverThreeCharactersCase("Mru")
	}
}

func ExampleReturnSuppliedNameIfOverThreeCharactersInitializationStatement() {
	fmt.Println(ReturnSuppliedNameIfOverThreeCharactersInitializationStatement("Mrunalini"))
	fmt.Println(ReturnSuppliedNameIfOverThreeCharactersInitializationStatement("Mru"))
	fmt.Println(ReturnSuppliedNameIfOverThreeCharactersInitializationStatement("Aaron"))
	// Output:
	// Mrunalini
	// Name must be over three characters
	// Aaron
}

func BenchmarkReturnSuppliedNameIfOverThreeCharactersInitializationStatement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnSuppliedNameIfOverThreeCharactersInitializationStatement("Mrunalini")
		ReturnSuppliedNameIfOverThreeCharactersInitializationStatement("Mru")
	}
}

func ExampleFizzBuzzCaseFallthrough() {
	fmt.Println(FizzBuzzCaseFallthrough(3))
	fmt.Println(FizzBuzzCaseFallthrough(10))
	fmt.Println(FizzBuzzCaseFallthrough(15))
	fmt.Println(FizzBuzzCaseFallthrough(14))
	fmt.Println(FizzBuzzCaseFallthrough(2))
	// Output:
	// Fizz Fallthrough
	// Buzz
	// FizzBuzz
	// Nothing
	//  Fallthrough
}
