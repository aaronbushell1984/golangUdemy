package recoverexample

import "fmt"

func ExampleCountRecoverPanic() {
	fmt.Println(CountRecoverPanic(0))
	// Output:
	// Calling Recursive...
	// Printing Number... : 0
	// Printing Number... : 1
	// Printing Number... : 2
	// Printing Number... : 3
	// Panicking, Value Over 3!
	// Defferred And Still captured... : 3
	// Defferred And Still captured... : 2
	// Defferred And Still captured... : 1
	// Defferred And Still captured... : 0
	// Recovered! Now Resuming...
}

func ExampleRecursivePrintUpToThree() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered...", r)
		}
	}()
	RecursivePrintUpToThree(0)
	// Output:
	// Printing Number... : 0
	// Printing Number... : 1
	// Printing Number... : 2
	// Printing Number... : 3
	// Panicking, Value Over 3!
	// Defferred And Still captured... : 3
	// Defferred And Still captured... : 2
	// Defferred And Still captured... : 1
	// Defferred And Still captured... : 0
	// Recovered... 4
}

func ExampleDemonstrateContinueAfterRecover() {
	fmt.Println(DemonstrateContinueAfterRecover(0))
	// Output:
	// Calling Recursive...
	// Printing Number... : 0
	// Printing Number... : 1
	// Printing Number... : 2
	// Printing Number... : 3
	// Panicking, Value Over 3!
	// Defferred And Still captured... : 3
	// Defferred And Still captured... : 2
	// Defferred And Still captured... : 1
	// Defferred And Still captured... : 0
	// Recovered! Now Resuming...
	// Continues After Recover...
}



