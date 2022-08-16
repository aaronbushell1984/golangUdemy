// Package recoverexample demonstrates recovering from panic in go
//
// Recover is only useful inside deferred functions
package recoverexample

import "fmt"

// RecursivePrintUpToThree will count from provided int until reaches four and will panic
func RecursivePrintUpToThree(i int) {
	if i > 3 {
		fmt.Println("Panicking, Value Over 3!")
		panic(fmt.Sprintf("%v", i))
	}
	fmt.Println("Printing Number... :", i)
	defer fmt.Println("Defferred And Still captured... :", i)
	RecursivePrintUpToThree(i + 1)
}

// CountRecoverPanic recovers a panic from RecursivePrintUpToThree:
//	defer func() {
//	if r := recover(); r != nil {
//		fmt.Println("Recovered! Now Resuming...")
//	}
//	}()
//	fmt.Println("Calling Recursive...")
//	RecursivePrintUpToThree(i)
//	return ""
func CountRecoverPanic(i int) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered! Now Resuming...")
		}
	}()
	fmt.Println("Calling Recursive...")
	RecursivePrintUpToThree(i)
	return ""
}

// DemonstrateContinueAfterRecover shows normal execution after a routine is recovered
func DemonstrateContinueAfterRecover(i int) string {
	CountRecoverPanic(i)
	return "Continues After Recover..."
}