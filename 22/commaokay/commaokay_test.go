package commaokay

import "fmt"

func ExampleConsumeMultipleReadOnlyChannels() {
	e, o, q := ConsumeMultipleReadOnlyChannels()
	for {
		select {
		case v := <-e:
			fmt.Printf("Even: %v\t", v)
		case v := <-o:
			fmt.Printf("Odd: %v\t", v)
		case v, ok := <-q:
			fmt.Printf("Quitting on: %v Function ran ok?: %v", v, ok)
			return
		}
	}
	// Output:
	// Even: 0	Odd: 1	Even: 2	Odd: 3	Even: 4	Odd: 5	Even: 6	Odd: 7	Even: 8	Odd: 9	Quitting on: 10 Function ran ok?: true
}

func ExampleConfirmValueReadFromChannel() {
	value, k := ConfirmValueReadFromChannel()
	fmt.Printf("%v %v", value, k)
	// Output:
	// 42 true
}

func ExampleConfirmValueReadFromChannelClose() {
	value, okay := ConfirmValueReadFromChannel()
	fmt.Printf("%v %v", value, okay)
	// Output:
	// 42 true
}
