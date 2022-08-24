package closure

import "fmt"

func ExampleIncrement() {
	firstCounter := Increment()
	secondCounter := Increment()
	fmt.Println(firstCounter())
	fmt.Println(firstCounter())
	fmt.Println(firstCounter())
	fmt.Println(secondCounter())
	fmt.Println(secondCounter())
	fmt.Println(firstCounter())
	fmt.Println(secondCounter())
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 4
	// 3
}
