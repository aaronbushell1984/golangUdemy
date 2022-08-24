package returnfunc

import "fmt"

func ExampleGetIntViaFunc() {
	fmt.Println(GetIntViaFunc()())
	// Output:
	// 42
}
