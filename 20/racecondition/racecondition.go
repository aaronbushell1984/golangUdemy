// Package racecondition demonstrates a race condition in go
package racecondition

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// RaceConditionCount creates 100 go routines all incrementing the same counter 100 times:
//	var counter int
// 	for i := 0; i < 100; i++ {
//		go func() {
// 			res := counter
//			runtime.Gosched()
//			res++
//			counter = res
//			wg.Done()
//		}()	
//	}
// This leads routines sharing the same memory space and a count of less than 100 in all cases
func RaceConditionCount() int {
	wg.Add(100)
	var counter int
	for i := 0; i < 100; i++ {
		go func() {
			res := counter
			runtime.Gosched()
			res++
			counter = res
			wg.Done()
		}()
		
	}
	wg.Wait()
	return counter
}