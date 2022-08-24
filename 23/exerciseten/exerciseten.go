// Package exerciseten is the tenth set of exercises from the course
package exerciseten

import "fmt"

// ReceiveFromChannel would block without the func literal - go routine to receive the value 42:
//	c := make(chan int)
//	go func() {
//		c <- 42
//	}()
//	return <-c
func ReceiveFromChannelFuncLiteral() int {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	return <-c
}

// ReceiveFromBufferedChannel would block without a buffer of exactly the right amout to receive the 42:
//	c := make(chan int, 1)
func ReceiveFromBufferedChannel() int {
	c := make(chan int, 1)
	c <- 42
	return <-c
}

// GetValueInReceiveOnlyChannel add 42 to a channel and returns it as receive only so:
//	<-GetValueInReceiveOnlyChannel()
// receives the value and cannot be misused by sending information to channel
func GetValueInReceiveOnlyChannel() <-chan int {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	return c
}

// Receive takes a receive only channel and ranges over the results to append to the resulting slice of int
func Receive(c <-chan int) []int {
	var numbers []int
	for v := range c {
		numbers = append(numbers, v)
	}
	return numbers
}

// Generate100 adds and returns numbers 0-99 in a receive only channel
//
// func literal and closing the channel on completion of loop prevents a block:
//	go func() {
//		for i := 0; i < 100; i++ {
//			c <- i
//		}
//		close(c)
//	}()
func Generate100() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

// QuitWithZero returns a receive only channel with zero to be used a quit channle later
func QuitWithZero() <-chan int {
	q := make(chan int)
	go func() {
		q <- 0
	}()
	return q
}

// ReceiveWithSelect loops indefinitely over first provided channal until the second sends in the quit value
//
// All values are added to a slice of int which returns on completion using either case:
//	var numbers []int
//	for {
//		select {
//		case v := <-c:
//			numbers = append(numbers, v)
//		case v := <-q:
//			numbers = append(numbers, v)
//		return numbers
//		}
//	}
func ReceiveWithSelect(c, q <-chan int) []int {
	var numbers []int
	for {
		select {
		case v := <-c:
			numbers = append(numbers, v)
		case v := <-q:
			numbers = append(numbers, v)
			return numbers
		}
	}
}

// GetOkFromOpenChannel shows ok is used to show a value is taken from channel successfully
func GetOkFromOpenChannel() bool {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	_, ok := <-c
	return ok
}

// GetOkFromClosedChannel shows that a channel which has had its value taken and is closed returns false into ok:
//	_, ok := <-c
//	close(c)
//	_, ok = <-c
func GetOkFromClosedChannel() string {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	_, ok := <-c
	fmt.Printf("The first ok is: %v\n", ok)
	close(c)
	_, ok = <-c
	return fmt.Sprintf("The second ok is: %v", ok)
}

// FanOut10GoRoutinesCountingToTen creates and returns 10 go routines all placing 1-10 into a read only channel
func FanOut10GoRoutinesCountingToTen() <-chan int {
	c := make(chan int)
	for i := 1; i < 11; i++ {
		go func() {
			for i := 1; i < 11; i++ {
				c <- i
			}
		}()
	}
	return c
}

// Fanin receives from the receive only channel 100 times and returns the values in a slice of int
func FanIn(out <-chan int) []int {
	var numbers []int
	for i := 0; i < 100; i++ {
		numbers = append(numbers, <-out)
	}
	return numbers
}
