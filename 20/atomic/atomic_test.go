package atomic

import (
	"fmt"
	"testing"
)

func ExampleRaceConditionCountWithAtomic() {
	fmt.Println(RaceConditionCountWithAtomic() == 100)
	// Output:
	// true
}

func BenchmarkRaceConditionCountWithAtomic(b *testing.B) {
	RaceConditionCountWithAtomic()
}
