package returnfunc

import "fmt"

func ExampleGetIntViaFunc() {
	theFunc := GetIntViaFunc()
	answer := theFunc()
	fmt.Println(answer)
	// Output:
	// 42
}