// Package rangechannel demonstrates ranging data in a channel until empty
package rangechannel

// SendReceiveClose returns the numbers 0-25 by looping into a channel:
//	go func() {
//	for i := 0; i < 25; i++ {
//		c <- i
//	}
//	close(c)
//	}()
// And ranging over the channel:
//	for v := range c {
//		res = append(res, v)
//	}
// The close is important here to not range over the channel and then meet a block:
//	close(c)
func SendReceiveClose() []int {
	c := make(chan int)
	var res []int
	go func() {
		for i := 0; i < 25; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		res = append(res, v)
	}
	return res
}

// ConsumeReadOnlyRange returns the numbers 0-25 by looping into a channel which returns a read only channel for consumption
func ConsumeReadOnlyRange() <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 25; i++ {
			c <- i
		}
	}()
	return c
}
