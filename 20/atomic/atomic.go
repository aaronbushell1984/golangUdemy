// Pacakge atomic demonstrates using atomic in go
//
// // Compare with racecondition package to see that condition solved here with atomic
package atomic

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// RaceConditionCountWithAtomic creates 100 go routines all incrementing the same counter 100 times:
//	var counter int
// 	for i := 0; i < 100; i++ {
//		go func() {
// 			atomic.AddInt64(&counter, 1)
//			runtime.Gosched()
//			wg.Done()
//		}()	
//	}
// This function is slower than a regular loop and faster than mutex:
//	BenchmarkRaceConditionCountWithAtomic-16 1000000000 0.0000880 ns/op 0 B/op 0 allocs/op
//	BenchmarkRaceConditionCountWithMutex-16	1000000000 0.0002359 ns/op 0 B/op 0	allocs/op
//	BenchmarkRegularCounterLoop-16		1000000000 0.0000003 ns/op 0 		allocs/op
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