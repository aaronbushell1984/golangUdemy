package racecondition

import "fmt"

func ExampleRaceConditionCount() {
	fmt.Println(RaceConditionCount() == 100)
	// Output:
	// false
}
