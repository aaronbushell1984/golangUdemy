package excercisetwo

import (
	"fmt"
	"testing"
)

func ExampleReturnDecimalBinaryHex() {
	fmt.Println(ReturnDecimalBinaryHex(10))
	fmt.Println(ReturnDecimalBinaryHex(32))
	fmt.Println(ReturnDecimalBinaryHex(100))
	// Output:
	// Decimal: 10 Binary: 1010 Hex: 0xa
	// Decimal: 32 Binary: 100000 Hex: 0x20
	// Decimal: 100 Binary: 1100100 Hex: 0x64
}

func ExampleCheckEquals() {
	fmt.Println(CheckEquals(1, 2))
	fmt.Println(CheckEquals(1, 1))
	fmt.Println(CheckEquals(-1, 1))
	fmt.Println(CheckEquals(0, 0))
	fmt.Println(CheckEquals(-10, -10))
	// Output:
	// false
	// true
	// false
	// true
	// true
}

func ExampleCheckLessOrEqual() {
	fmt.Println(CheckLessOrEqual(1, 2))
	fmt.Println(CheckLessOrEqual(1, 1))
	fmt.Println(CheckLessOrEqual(-1, 1))
	fmt.Println(CheckLessOrEqual(0, 0))
	fmt.Println(CheckLessOrEqual(-10, -10))
	fmt.Println(CheckLessOrEqual(10, -10))
	// Output:
	// true
	// true
	// true
	// true
	// true
	// false
}

func ExampleCheckGreaterOrEqual() {
	fmt.Println(CheckGreaterOrEqual(1, 2))
	fmt.Println(CheckGreaterOrEqual(1, 1))
	fmt.Println(CheckGreaterOrEqual(-1, 1))
	fmt.Println(CheckGreaterOrEqual(0, 0))
	fmt.Println(CheckGreaterOrEqual(-10, -10))
	fmt.Println(CheckGreaterOrEqual(10, -10))
	// Output:
	// false
	// true
	// false
	// true
	// true
	// true
}

func ExampleCheckNotEqual() {
	fmt.Println(CheckNotEqual(1, 2))
	fmt.Println(CheckNotEqual(1, 1))
	fmt.Println(CheckNotEqual(-1, 1))
	fmt.Println(CheckNotEqual(0, 0))
	fmt.Println(CheckNotEqual(-10, -10))
	fmt.Println(CheckNotEqual(10, -10))
	// Output:
	// true
	// false
	// true
	// false
	// false
	// true
}

func ExampleCheckLessThan() {
	fmt.Println(CheckLessThan(1, 2))
	fmt.Println(CheckLessThan(1, 1))
	fmt.Println(CheckLessThan(-1, 1))
	fmt.Println(CheckLessThan(0, 0))
	fmt.Println(CheckLessThan(-10, -10))
	fmt.Println(CheckLessThan(10, -10))
	// Output:
	// true
	// false
	// true
	// false
	// false
	// false
}

func ExampleCheckGreaterThan() {
	fmt.Println(CheckGreaterThan(1, 2))
	fmt.Println(CheckGreaterThan(1, 1))
	fmt.Println(CheckGreaterThan(-1, 1))
	fmt.Println(CheckGreaterThan(0, 0))
	fmt.Println(CheckGreaterThan(-10, -10))
	fmt.Println(CheckGreaterThan(10, -10))
	// Output:
	// false
	// false
	// false
	// false
	// false
	// true
}

func TestReturnConstTrafficLights(t *testing.T) {
	want := "RED, AMBER, GREEN"
	if got := ReturnConstTrafficLights(); got != want {
		t.Errorf("ReturnConstTrafficLights: Got: %q, Want: %q", got, want)
	}
}
