// Package mutex demonstrates using mutex in go
//
// Compare with racecondition package to see that condition solved here with mutex
package mutex

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

// RaceConditionCountWithMutex creates 100 go routines all incrementing the same counter 100 times:
//	var counter int
// 	for i := 0; i < 100; i++ {
//		go func() {
//			mu.Lock()
// 			res := counter
//			runtime.Gosched()
//			res++
//			counter = res
//			mu.Unlock()
//			wg.Done()
//		}()	
//	}
// The memory is Locked by:
//	mu.Lock()
// And Unlocked by:
//	mu.Unlock()
// Meaning each routine locks the space if lucky enough to gain access before unlocking after updating the count
func RaceConditionCountWithMutex() int {
	wg.Add(100)
	var counter int
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			res := counter
			runtime.Gosched()
			res++
			counter = res
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}

// RegularCounterLoop increments and returns a counter for benchmarking against RaceConditionCountWithMutex
//
// A regular loop is much faster:
//	BenchmarkRaceConditionCountWithMutex-16	1000000000 0.0002359 ns/op 0 B/op 0	allocs/op
//	BenchmarkRegularCounterLoop-16		1000000000 0.0000003 ns/op 0 		allocs/op
func RegularCounterLoop() int {
	var counter int
	for i := 0; i < 100; i++ {
		res := counter
		res++
		counter = res
	}
	return counter
}