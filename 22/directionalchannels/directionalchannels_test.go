package directionalchannels

import "fmt"

func ExampleUseSendAndReceiveOnlyChannels() {
	n := 42
	fmt.Println(UseSendAndReceiveOnlyChannels(n))
	// Output:
	// 84
}

func ExampleConsumeReadOnlyChannel() {
	fmt.Println(<-ConsumeReadOnlyChannel())
	// Output:
	// Data I want you to Receive and not Write - and then close channel
}
