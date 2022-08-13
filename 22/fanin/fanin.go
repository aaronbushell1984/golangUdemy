// Package fanin demonstrates the fan in concurrency design pattern in go
//
// This is used after launching nultiple routines to collect the data into a single channel for results
package fanin

import (
	"fmt"
	"sync"
)

// SendOddAndEvenToChannel sends numbers 1-10 to respective even and odd channels
func SendToOddAndEvenChannels(even, odd  chan<- string) {
	for i := 1; i < 11; i++ {
		if i % 2 == 0 {
			even <- fmt.Sprintf("Even: %v", i)
		}
		if i % 2 != 0 {
			odd <- fmt.Sprintf("Odd: %v", i)
		}
	}
	defer close(odd)
	defer close(even)
}

// FanInReceiveOddAndEvenChannel takes two channels and sends in to one channel (fan in)
func FanInReceiveOddAndEvenChannel(even, odd <-chan string, fanin chan<- string) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}

// SendBoringChannel sends a message to a receive channel using the provided name
func SendBoringChannel(name string) <-chan string {
	c := make(chan string)
	go func(){
		defer close(c)
		for i := 0; i < 3; i++ {
			c <- fmt.Sprintf("%v %v is boring", i, name)
		}
	}()
	return c
}

// FanInToChannel takes in two receive channels and returns a receive channel
func FanInToChannel(ann, bob <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-ann
		}
	}()
	go func() {
		for {
			c <- <-bob
		}
	}()
	return c
}