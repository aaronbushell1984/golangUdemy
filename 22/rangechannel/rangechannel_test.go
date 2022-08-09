package rangechannel

import "fmt"

func ExampleSendReceiveClose() {
	fmt.Println(SendReceiveClose())
	// Output:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24]
}

func ExampleConsumeReadOnlyRange() {
	in := ConsumeReadOnlyRange()
	var res []int
	for v := range in {
		res = append(res, v)
	}
	fmt.Println(res)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24]

}


