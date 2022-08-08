// Package channnels demonstrates the use of channels in go
//
// Channels BLOCK for example:
//	func ChannelBlock() int {
//		c:= make(chan int)
//		c <- 42
//		return <-c
//	}
// will run forever if called as:
//	c <- 42
// will only execute if a process can receive 42 at SAME time
//
// Channels can be launched safely with a go routine:
//	go func() {
//		c <- 42
//	}()
// Or a buffer which is large enough to receive the values required:
//	c := make(chan int, 1)
//
// Channels will be blocked if the buffer is too small:
//	c := make(chan int, 1)
//	c <- 42
//	c <- 43
//	return <-c
package channels

import "fmt"

// ChannelBlock demonstrates a channel blocks
func ChannelBlock() int {
	c:= make(chan int)
	c <- 42
	return <-c
}

// GetChannelType returns the type of any argument and is used to demonstrate the type of a channel in example
func GetChannelType(channel interface{}) string {
	return fmt.Sprintf("%T", channel)
}

// ChannelBlockedCodeContinues demonstrates code runs with channel blocked/waiting in a go routine:
//	go func() {
//		c <- 42
//	}()
// Value is received into return from channel before function exits:
//	return <-c
func ChannelBlockedCodeContinues() int {
	c:= make(chan int)
	go func() {
		c <- 42
	}()
	return <-c
}

// BufferedChannelCodeContinues because a buffer argument of 1 is passed:
//	c := make(chan int, 1)
// Exactly 1 channel is passed a value:
//	c <- 42
// So there is never more than one in the buffer and code runs
func BufferedChannelCodeContinues() int {
	c := make(chan int, 1)
	c <- 42
	return <-c
}

//  BufferedChannelTooSmallCodeBlocked demonstrates a buffer which is too small blocks
func BufferedChannelTooSmallCodeBlocked() int {
	c := make(chan int, 1)
	c <- 42
	c <- 43
	return <-c
}