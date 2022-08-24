package callback

import "fmt"

func ExampleGetSum() {
	fmt.Println(GetSum(1, 2, 3))
	fmt.Println(GetSum(-1, 2, 3))
	// Output:
	// 6
	// 4
}

func ExampleGetEven() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers0 := []int{0, 2, 3, 4, 5}
	numbersMinus := []int{-1, -2, -3, -4, -5}
	fmt.Println(GetEven(numbers...))
	fmt.Println(GetEven(numbers0...))
	fmt.Println(GetEven(numbersMinus...))
	// Output:
	// [2 4]
	// [2 4]
	// [-2 -4]
}

func ExampleGetEvenSum() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers0 := []int{0, 2, 3, 4, 5}
	numbersMinus := []int{-1, -2, -3, -4, -5}
	fmt.Println(GetEvenSum(GetSum, numbers...))
	fmt.Println(GetEvenSum(GetSum, numbers0...))
	fmt.Println(GetEvenSum(GetSum, numbersMinus...))
	// Output:
	// 6
	// 6
	// -6
}

func ExampleGetOdd() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers0 := []int{0, 2, 3, 4, 5}
	numbersMinus := []int{-1, -2, -3, -4, -5}
	fmt.Println(GetOdd(numbers...))
	fmt.Println(GetOdd(numbers0...))
	fmt.Println(GetOdd(numbersMinus...))
	// Output:
	// [1 3 5]
	// [3 5]
	// [-1 -3 -5]
}

func ExampleGetOddSum() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers0 := []int{0, 2, 3, 4, 5}
	numbersMinus := []int{-1, -2, -3, -4, -5}
	fmt.Println(GetOddSum(numbers...))
	fmt.Println(GetOddSum(numbers0...))
	fmt.Println(GetOddSum(numbersMinus...))
	// Output:
	// 9
	// 8
	// -9
}
