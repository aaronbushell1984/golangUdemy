package mutex

import (
	"fmt"
	"testing"
)

func ExampleRaceConditionCountWithMutex() {
	fmt.Println(RaceConditionCountWithMutex() == 100)
	// Output:
	// true
}

func ExampleRegularCounterLoop() {
	fmt.Print(RegularCounterLoop() == 100)
	// Output:
	// true
}

func BenchmarkRaceConditionCountWithMutex(b *testing.B) {
	RaceConditionCountWithMutex()
}

func BenchmarkRegularCounterLoop(b *testing.B) {
	RegularCounterLoop()
}
