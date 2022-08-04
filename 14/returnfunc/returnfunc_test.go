package returnfunc

import "fmt"

func ExampleGetIntViaFunc() {
	theFunc := GetIntViaFunc()
	fmt.Println(theFunc())
	// Output:
	// 42
}