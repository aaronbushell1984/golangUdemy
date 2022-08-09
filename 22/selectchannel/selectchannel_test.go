package selectchannel

import "fmt"

func ExampleConsumeMultipleReadOnlyChannels() {
	e, o, q := ConsumeMultipleReadOnlyChannels()
	for {
		select {
		case v := <-e:
			fmt.Printf("Even: %v\t", v)
		case v := <-o:
			fmt.Printf("Odd: %v\t", v)
		case v := <-q:
			fmt.Printf("Quitting on: %v", v)
			return
		}		
	}
	// Output:
	// Even: 0	Odd: 1	Even: 2	Odd: 3	Even: 4	Odd: 5	Even: 6	Odd: 7	Even: 8	Odd: 9	Quitting on: 10
}
