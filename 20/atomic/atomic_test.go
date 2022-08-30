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
	for i := 0; i < b.N; i++ {
		RaceConditionCountWithAtomic()
	}
}
