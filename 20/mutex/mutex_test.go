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
	for i := 0; i < b.N; i++ {
		RaceConditionCountWithMutex()
	}
}

func BenchmarkRegularCounterLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RegularCounterLoop()
	}
}
