// Package commaokay demonstrates the use of comma okay idiom using channels
package commaokay

// ConsumeMultipleReadOnlyChannels performs the same as that in selectchannel
//
// The example catches ok in the final case in select statement:
//	case v, ok := <-q:
//		fmt.Printf("Quitting on: %v Function ran ok?: %v", v, ok)
// As this case is run a true is returned to ok
func ConsumeMultipleReadOnlyChannels() (<-chan int, <-chan int, <-chan int) {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				even <- i
			}
			if i%2 != 0 {
				odd <- i
			}
		}
		quit <- 10
	}()
	return even, odd, quit
}

// ConfirmValueReadFromChannel allows consumption of 42 sent:
//		go func() {
//			c <- 42
//		}()
// ok is used to check channel receives 42 successfully
func ConfirmValueReadFromChannel() (int, bool) {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	v, ok := <-c
	return v, ok
}

// ConfirmValueReadFromChannelClose replicates ConfirmValueReadFromChannel but also closes channel
func ConfirmValueReadFromChannelClose() (int, bool) {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 42
	}()
	v, ok := <-c
	return v, ok
}
