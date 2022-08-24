package channels

import "fmt"

// THIS TEST NEVER RECEIVES AN OUTPUT AS CHANNELS BLOCK
// func ExampleChannelBlock() {
// 	fmt.Println(ChannelBlock())
// 	// Output:
// 	//
// }

func ExampleGetChannelType() {
	c := make(chan int)
	fmt.Println(GetChannelType(c))
	// Output:
	// chan int
}

func ExampleChannelBlockedCodeContinues() {
	fmt.Println(ChannelBlockedCodeContinues())
	// Output:
	// 42
}

func ExampleBufferedChannelCodeContinues() {
	fmt.Println(BufferedChannelCodeContinues())
	// Output:
	// 42
}

// THIS TEST NEVER RECEIVES AN OUTPUT AS CHANNELS BLOCK DUE TO BUFFER BEING TOO SMALL
// func ExampleBufferedChannelTooSmallCodeBlocked() {
// 	fmt.Println(BufferedChannelTooSmallCodeBlocked())
// 	// Output:
// 	//
// }
