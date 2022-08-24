// Package directional channels demonstrate one way channels in go
//
// Channels can be two way:
// c := make (chan int)
// Send only:
//	c := make(chan<- int)
// Or receive only:
//	c := make(<-chan int)
// A channel can be converted to a more restrictive channel
//
// This may be in argument or return; A 2 way channel returning a read only channel:
//	func() <-chan int {
//		c := make(chan int)
//		return c
//	}
package directionalchannels

// SendOnlyChannel will only send an int to a channel
func SendOnlyChannel(c chan<- int, num int) {
	c <- num
}

// ReceiveOnlyChannel will only receive the result from a channel
func ReceiveOnlyChannel(c <-chan int) int {
	return <-c
}

// UseSendAndReceiveOnlyChannels combines a read only and send only channel
//
// Creating a 2 way channel:
//	c := make(chan int)
// Launching go routines which will only send:
//	go SendOnlyChannel(c, num)
// Receiving channels and reading result for return:
//	one := ReceiveOnlyChannel(c)
func UseSendAndReceiveOnlyChannels(num int) int {
	c := make(chan int)
	go SendOnlyChannel(c, num)
	go SendOnlyChannel(c, num)
	one := ReceiveOnlyChannel(c)
	two := ReceiveOnlyChannel(c)
	return one + two
}

// ConsumeReadOnlyChannel allows consumption of a string without being able to write to channel:
//
// A value of two way channel type is created:
//	c := make(chan string)
// A go routine is started:
//	go func() {...}()
// The channel is closed after data is read using defer:
//	go func() {
//		defer close(c)
//		readOnly := "Data I want you to Receive and not Write - and then close channel"
//		c <- readOnly
//	}()
// The channel is converted to read only by returning:
//	return c
// as a read only channel:
//	func ConsumeReadOnlyChannel() <-chan string {
// A consumer may consume this information by reading from the channel:
//	fmt.Println(<-ConsumeReadOnlyChannel())
func ConsumeReadOnlyChannel() <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		readOnly := "Data I want you to Receive and not Write - and then close channel"
		c <- readOnly
	}()
	return c
}
