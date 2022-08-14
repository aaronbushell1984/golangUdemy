// Package fanout demonstrates fanning out of channels in go
//
// Typically used to run the same task repetitively. Such as encoding a video
//
// This allows multiple cores to execute in parallel if available
package fanout

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Populate creates a number of channels which receive their index starting at 1nnel and closes it on completion
//
// This can be launched inside a go routine and results ranged over as per example
func Populate(c chan int, count int) {
	for i := 0; i < count; i++ {
		c <- i
	}
	close(c)
}

// TimeConsumingWork takes in a value and returns it after a random wait of up to the value in microseconds
//
// This is used to simulate work for go routines to complete with FanOutIn
func TimeConsumingWork(value int) int {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(value)
	time.Sleep(time.Microsecond * time.Duration(n))
	return value
}

// FanOutIn takes two channels:
//	(numberPopulateChannel, receiveChannel chan int)
// The first to creates a number of channels using Populate
//
// The second receives the output of TimeConsumingWork:
//	receiveChannel <- TimeConsumingWork(v2)
// For each channel previously populated a new channel is launched to receive this work:
//	for v := range numberPopulateChannel {
// 		wg.Add(1)
// 		go func(v2 int) {
//			receiveChannel <- TimeConsumingWork(v2)
//			wg.Done()
//		}(v)
//	}
func FanOutIn(numberPopulateChannel, receiveChannel chan int) {
	var wg sync.WaitGroup
	for v := range numberPopulateChannel {
		wg.Add(1)
		go func(v2 int) {
			receiveChannel <- TimeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(receiveChannel)
}

// ConcurrentTimeConsumingWork combines Populate, TimeConsumingWork and FanInOut to launch "routines" number of TimeConsumingWork
func ConcurrentTimeConsumingWork(routines int) []int {
	var numbers []int
	numberPopulateChannel := make(chan int)
	receiveChannel := make(chan int)
	go Populate(numberPopulateChannel, routines)
	go FanOutIn(numberPopulateChannel, receiveChannel)
	for v := range receiveChannel {
		numbers = append(numbers , v)
	}
	sort.Ints(numbers)
	return numbers
}

// SequentialTimeConsumingWork is to compare the same task as ConcurrentTimeConsumingWork in a non concurrent pattern
//
// Output of benchmark tests on a 16 core processor show concurrency is faster:
//	BenchmarkConcurrentTimeConsumingWork-16         1000000000               0.001441 ns/op
//	ok      /golangUdemy/22/fanout       0.019s
//	BenchmarkSequentialTimeConsumingWork-16         1000000000               0.08071 ns/op
//	ok      /golangUdemy/22/fanout       0.745s
func SequentialTimeConsumingWork(count int) []int {
	var numbers []int
	for i := 1; i < count; i++ {
		numbers = append(numbers , TimeConsumingWork(i))
	}
	return numbers
}

// ThrottleFanOutIn allows number of fan out routines to be restricted to the throttle
//
// Compare with FanOutIn
func ThrottleFanOutIn(numberPopulateChannel chan int, receiveChannel chan int, throttle int) {
	var wg sync.WaitGroup
	wg.Add(throttle)
	for i := 0; i < throttle; i++ {
		go func() {
			for v := range numberPopulateChannel {	
				func(v2 int) {
					receiveChannel <- TimeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(receiveChannel)
}

// ThrottleConcurrentTimeConsumingWork is to compare with ConcurrentTimeConsumingWork when routines are throttled
//
// Benchmark tests show that reducing the amount of routines for this process can be faster:
//	BenchmarkConcurrentTimeConsumingWork-16         1000000000               0.001441 ns/op
//	ok      /golangUdemy/22/fanout       0.019s
//	BenchmarkThrottleConcurrentTimeConsumingWork-16         1000000000               0.001249 ns/op
//	ok      /golangUdemy/22/fanout       0.015s
func ThrottleConcurrentTimeConsumingWork(count int, throttle int) []int {
	var numbers []int
	numberPopulateChannel := make(chan int)
	receiveChannel := make(chan int)
	go Populate(numberPopulateChannel, count)
	go ThrottleFanOutIn(numberPopulateChannel, receiveChannel, throttle)
	for v := range receiveChannel {
		numbers = append(numbers , v)
	}
	sort.Ints(numbers)
	return numbers
}
