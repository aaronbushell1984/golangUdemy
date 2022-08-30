// Package atomic demonstrates using atomic in go
//
// Compare with race condition package to see that condition solved here with atomic
package atomic

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// RaceConditionCountWithAtomic creates 100 go routines all incrementing the same counter 100 times:
//
//	var counter int
//	for i := 0; i < 100; i++ {
//		go func() {
//			atomic.AddInt64(&counter, 1)
//			runtime.Gosched()
//			wg.Done()
//		}()
//	}
//
// This function is slower than a regular loop and faster than mutex:
//
//	BenchmarkRaceConditionCountWithAtomic-8 18980 61536 ns/op
//	BenchmarkRaceConditionCountWithMutex-8	9008 131830 ns/op
//	BenchmarkRegularCounterLoop-8		25266441 46.43 ns/op
func RaceConditionCountWithAtomic() int64 {
	wg.Add(100)
	var counter int64
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()

	}
	wg.Wait()
	return counter
}
